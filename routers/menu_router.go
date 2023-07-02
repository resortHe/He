package routers

import "go_vue/api"

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("menus", menuApi.MenuCreateView)
	router.GET("menus", menuApi.MenuListView)
	router.GET("menus_name_list", menuApi.MenuNameList)
	router.PUT("menus/:id", menuApi.MenuUpdateView)
	router.DELETE("menus", menuApi.MenuRemoveView)
	router.GET("menus/:id", menuApi.MenuDetailView)
}
