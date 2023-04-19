package initial

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/public"
	"github.com/jiaruling/golang_utils/lib"
)

func init() {
	// 初始化项目目录
	lib.InitDir([]string{global.LogFilePath, global.FILEPATH}, []string{"log"})
	// 加载配置文件
	initConfigFile("./conf/dev")
	initConfEnv("./conf/dev")
	// 初始化日志配置
	lib.InitLog(global.ConfigBase.Log.LogFileDir, global.ConfigBase.Log.AppName,
		global.ConfigBase.Log.MaxSize, global.ConfigBase.Log.MaxBackups,
		global.ConfigBase.Log.MaxAge, global.ConfigBase.Log.Dev,
	)
	// 初始化数据库连接
	trace, log := public.GetTraceAndLog()
	log.Error(trace, lib.DLTagMySqlFailed, map[string]interface{}{"hint": "连接数据库失败", "hint_code": 1001})
	log.Warn(trace, lib.DLTagRedisFailed, map[string]interface{}{"hint": "连接REDIS失败", "hint_code": 1001})
}

func initConfigFile(path string) {
	if err := lib.ParseConfig(filepath.Join(path, "base.toml"), &global.ConfigBase, false); err != nil {
		fmt.Println("加载配置文件失败001: ", err.Error())
		os.Exit(1)
	}
	if err := lib.ParseConfig(filepath.Join(path, "mysql.toml"), &global.ConfigMySQL, false); err != nil {
		fmt.Println("加载配置文件失败002: ", err.Error())
		os.Exit(1)
	}
	if err := lib.ParseConfig(filepath.Join(path, "redis.toml"), &global.ConfigRedis, false); err != nil {
		fmt.Println("加载配置文件失败003: ", err.Error())
		os.Exit(1)
	}
	if err := lib.ParseConfigViper(filepath.Join(path, "base.toml"), "toml", false); err != nil {
		fmt.Println("加载配置文件失败004: ", err.Error())
		os.Exit(1)
	}
	return
}

func initConfEnv(path string) {
	p := strings.Split(path, "/")
	global.ConfEnv = p[len(p)-1]
	return
}
