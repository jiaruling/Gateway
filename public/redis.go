package public

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jiaruling/golang_utils/lib"
)

func RedisConfPipline(pip ...func(c redis.Conn)) error {
	c := lib.GetRedis()
	defer c.Close()
	for _, f := range pip {
		f(c)
	}
	return nil
}

func RedisConfDo(commandName string, args ...interface{}) (interface{}, error) {
	c := lib.GetRedis()
	defer c.Close()
	return c.Do(commandName, args...)
}
