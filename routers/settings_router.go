package routers

import (
	"go_vue/api"
)

// SettingsRoute 系统路由
func (router RouterGroup) SettingsRoute() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)

}
