package public

import (
	"sync"

	"golang.org/x/time/rate"
)

// 限流器

var FlowLimiterHandler *FlowLimiter

type FlowLimiter struct {
	FlowLmiterMap   map[string]*FlowLimiterItem
	FlowLmiterSlice []*FlowLimiterItem
	Locker          sync.RWMutex
}

type FlowLimiterItem struct {
	ServiceName string
	Limter      *rate.Limiter
}

func NewFlowLimiter() *FlowLimiter {
	return &FlowLimiter{
		FlowLmiterMap:   map[string]*FlowLimiterItem{},
		FlowLmiterSlice: []*FlowLimiterItem{},
		Locker:          sync.RWMutex{},
	}
}

func init() {
	FlowLimiterHandler = NewFlowLimiter()
}

// 获取限流器
func (counter *FlowLimiter) GetLimiter(serverName string, qps float64) (*rate.Limiter, error) {
	// 存在限流器直接返回
	for _, item := range counter.FlowLmiterSlice {
		if item.ServiceName == serverName {
			return item.Limter, nil
		}
	}

	newLimiter := rate.NewLimiter(rate.Limit(qps), int(qps*3))
	item := &FlowLimiterItem{
		ServiceName: serverName,
		Limter:      newLimiter,
	}

	// 加入slice和map
	counter.FlowLmiterSlice = append(counter.FlowLmiterSlice, item)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	counter.FlowLmiterMap[serverName] = item
	return newLimiter, nil
}

// 重置限流器
func (counter *FlowLimiter) ResetLimiter(serverName string, qps float64) {
	if _, ok := counter.FlowLmiterMap[serverName]; ok {
		newLimiter := rate.NewLimiter(rate.Limit(qps), int(qps*3))
		item := &FlowLimiterItem{
			ServiceName: serverName,
			Limter:      newLimiter,
		}
		counter.FlowLmiterMap[serverName] = item
		for i := range counter.FlowLmiterSlice {
			if counter.FlowLmiterSlice[i].ServiceName == serverName {
				counter.FlowLmiterSlice[i] = item
			}
		}
	}
}
