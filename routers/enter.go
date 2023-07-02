package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go_vue/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiRouterGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统设置api
	routerGroupApp.SettingsRoute()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertRouter()
	routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.TagRouter()
	return r
}
