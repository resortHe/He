package routers

import "go_vue/api"

// ImagesRouter 图片路由
func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.GET("images", imagesApi.ImageListView)
	router.GET("images_name_list", imagesApi.ImageNameListView)
	router.POST("images", imagesApi.ImageUploadView)
	router.DELETE("images", imagesApi.ImageRemoveView)
	router.PUT("images", imagesApi.ImageUpdateView)
}
