package main

import (
	"go_vue/core"
	"go_vue/global"
	"go_vue/routers"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	router := routers.InitRouter()
	addr := global.Config.System.Add()
	global.Log.Infof("go_db运行在：%s", addr)
	router.Run(addr)
}
