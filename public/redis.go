package public

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jiaruling/golang_utils/lib"
)

func RedisConfPipline(pip ...func(c *redis.Client)) error {
	c := lib.GetRedis()
	for _, f := range pip {
		f(c)
	}
	return nil
}

func RedisConfDo(commandName string, args ...interface{}) (int64, error) {
	rdb := lib.GetRedis()
	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val, 10, 64)
}
