package public

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/golang_utils/lib"
)

type RedisFlowCountService struct {
	AppID       string
	Interval    time.Duration
	QPS         int64
	Unix        int64
	TickerCount int64
	TotalCount  int64
}

func NewRedisFlowCountService(appID string, interval time.Duration) *RedisFlowCountService {
	reqCounter := &RedisFlowCountService{
		AppID:    appID,
		Interval: interval,
		QPS:      0,
		Unix:     0,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			tickerCount := atomic.LoadInt64(&reqCounter.TickerCount) //获取数据
			atomic.StoreInt64(&reqCounter.TickerCount, 0)            //重置数据

			currentTime := time.Now()
			dayKey := reqCounter.GetDayKey(currentTime)
			hourKey := reqCounter.GetHourKey(currentTime)
			if err := RedisConfPipline(func(c *redis.Client) {
				c.HSet(context.Background(), "INCRBY", dayKey, tickerCount, 86400*2)
				c.HSet(context.Background(), "EXPIRE", dayKey, tickerCount, 86400*2)
				c.HSet(context.Background(), "INCRBY", hourKey, tickerCount, 86400*2)
				c.HSet(context.Background(), "EXPIRE", hourKey, tickerCount, 86400*2)
			}); err != nil {
				fmt.Println("RedisConfPipline err", err)
				continue
			}

			totalCount, err := reqCounter.GetDayData(currentTime)
			if err != nil {
				fmt.Println("reqCounter.GetDayData err", err)
				continue
			}
			nowUnix := time.Now().Unix()
			if reqCounter.Unix == 0 {
				reqCounter.Unix = time.Now().Unix()
				continue
			}
			tickerCount = totalCount - reqCounter.TotalCount
			if nowUnix > reqCounter.Unix {
				reqCounter.TotalCount = totalCount
				reqCounter.QPS = tickerCount / (nowUnix - reqCounter.Unix)
				reqCounter.Unix = time.Now().Unix()
			}
		}
	}()
	return reqCounter
}

func (o *RedisFlowCountService) GetDayKey(t time.Time) string {
	dayStr := t.In(lib.TimeLocation).Format("20060102")
	return fmt.Sprintf("%s_%s_%s", global.RedisFlowDayKey, dayStr, o.AppID)
}

func (o *RedisFlowCountService) GetHourKey(t time.Time) string {
	hourStr := t.In(lib.TimeLocation).Format("2006010215")
	return fmt.Sprintf("%s_%s_%s", global.RedisFlowHourKey, hourStr, o.AppID)
}

func (o *RedisFlowCountService) GetHourData(t time.Time) (int64, error) {
	return RedisConfDo("GET", o.GetHourKey(t))
}

func (o *RedisFlowCountService) GetDayData(t time.Time) (int64, error) {
	return RedisConfDo("GET", o.GetDayKey(t))
}

// 原子增加
func (o *RedisFlowCountService) Increase() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		atomic.AddInt64(&o.TickerCount, 1)
	}()
}
