package public

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jiaruling/Gateway/global"
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
	go func(rc *RedisFlowCountService) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			tickerCount := atomic.LoadInt64(&rc.TickerCount) //获取数据
			atomic.StoreInt64(&rc.TickerCount, 0)            //重置数据

			currentTime := time.Now()
			dayKey := rc.GetDayKey(currentTime)   // 获取天的key
			hourKey := rc.GetHourKey(currentTime) // 获取当前小时的key
			// 持久化到redis
			RedisConfPipline(func(c *redis.Client) {
				ctx := context.Background()
				c.IncrBy(ctx, dayKey, tickerCount)       // 累加流量
				c.Expire(ctx, dayKey, global.DataAlive)  // 设置过期时间
				c.IncrBy(ctx, hourKey, tickerCount)      // 累加流量
				c.Expire(ctx, hourKey, global.DataAlive) // 设置过期时间
			})
			// 获取总流量
			totalCount, err := rc.GetDayData(currentTime)
			if err != nil {
				fmt.Println("reqCounter.GetDayData err", err)
				continue
			}
			nowUnix := time.Now().Unix()
			if rc.Unix == 0 {
				rc.Unix = time.Now().Unix()
				continue
			}
			// 计算QPS
			tickerCount = totalCount - rc.TotalCount
			if nowUnix > rc.Unix {
				rc.TotalCount = totalCount
				rc.QPS = tickerCount / (nowUnix - rc.Unix)
				rc.Unix = time.Now().Unix()
			}
		}
	}(reqCounter)
	return reqCounter
}

func (o *RedisFlowCountService) GetDayKey(t time.Time) string {
	dayStr := t.Format("20060102")
	return fmt.Sprintf("%s_%s_%s", global.RedisFlowDayKey, dayStr, o.AppID)
}

func (o *RedisFlowCountService) GetHourKey(t time.Time) string {
	hourStr := t.Format("2006010215")
	return fmt.Sprintf("%s_%s_%s", global.RedisFlowHourKey, hourStr, o.AppID)
}

func (o *RedisFlowCountService) GetHourData(t time.Time) (int64, error) {
	return RedisConfDo(o.GetHourKey(t))
}

func (o *RedisFlowCountService) GetDayData(t time.Time) (int64, error) {
	return RedisConfDo(o.GetDayKey(t))
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
