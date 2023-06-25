package routers

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	apiRouterGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统设置api
	routerGroupApp.SettingsRoute()
	return r
}
