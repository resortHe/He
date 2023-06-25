package api

import "go_vue/api/settings_api"

type AipGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(AipGroup)
