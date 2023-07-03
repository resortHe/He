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
		group.GET("count", messageApi.MessageListView)
		group.POST("record", messageApi.MessageRecordView)
		group.DELETE("remove", messageApi.MessageRemoveView)
	}
	router.GET("messages_all", middleware.JwtAdmin(), messageApi.MessageListAllView)
}

