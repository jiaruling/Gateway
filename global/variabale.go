package global

import "github.com/jiaruling/Gateway/global/conf"

var (
	ConfigBase  conf.BaseConf
	ConfigMySQL conf.MySQLConf
	ConfigRedis conf.RedisConf
	ConfEnv     string
	LoadTypeMap = map[int]string{
		LoadTypeHTTP: "HTTP",
		LoadTypeTCP:  "TCP",
		LoadTypeGRPC: "GRPC",
	}
)
