package routers

import "go_vue/api"

// AdvertRouter 广告路由
func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertApi
	router.POST("advert", advertApi.AdvertCreateView)
	router.GET("advert", advertApi.AdvertListView)
	router.PUT("advert", advertApi.AdvertUpdateView)
	router.DELETE("advert", advertApi.AdvertRemoveView)
}
