package routers

import "go_vue/api"

func (router RouterGroup) TagRouter() {
	tagtApi := api.ApiGroupApp.TagApi
	router.POST("tags", tagtApi.TagCreateView)
	router.GET("tags", tagtApi.TagListView)
	router.PUT("tags/:id", tagtApi.TagUpdateView)
	router.DELETE("tags", tagtApi.TagRemoveView)
}
