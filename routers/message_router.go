package routers

import (
	"go_vue/api"
	"go_vue/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	group := router.Group("message").Use(middleware.JwtAuth())
	{
		group.POST("create", messageApi.MessageCreateView)
	}

}
