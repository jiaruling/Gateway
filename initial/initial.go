package initial

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jiaruling/Gateway/global"
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
	initDB()
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

func initDB() {
	// Mysql
	for _, m := range global.ConfigMySQL.List {
		mysql := lib.NewMysqlGorm(m.User, m.Password, m.Ip, m.Port, m.DB)
		mysql.Name = m.Name
		mysql.Parameter = m.Parameter
		mysql.InitMysqlGorm()
	}
	// Redis
	for _, r := range global.ConfigRedis.List {
		redis := lib.NewRedis(fmt.Sprintf("%s:%v", r.Ip, r.Port), r.Password, r.DB)
		redis.Name = r.Name
		redis.InitRedis()
	}
}
