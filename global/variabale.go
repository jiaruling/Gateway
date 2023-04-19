package global

import "github.com/jiaruling/Gateway/global/conf"

var (
	ConfigBase  conf.BaseConf
	ConfigMySQL conf.MySQLConf
	ConfigRedis conf.RedisConf
	ConfEnv     string
)
