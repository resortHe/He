package routers

import (
	"go_vue/api"
)

func (router RouterGroup) SettingsRoute() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)
}
