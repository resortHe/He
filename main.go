package main

import (
	"go_vue/core"
	_ "go_vue/docs"
	"go_vue/flag"
	"go_vue/global"
	"go_vue/routers"
)

// @title go_vue API文档
// @version 1.0
// @description  go_vue API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	//连接Redis
	global.Redis = core.ConnectRedis()
	//命令行绑定参数
	parse := flag.Parse()
	if flag.IsWebStop(parse) {
		flag.SwitchOption(parse)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Add()
	global.Log.Infof("go_db运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
