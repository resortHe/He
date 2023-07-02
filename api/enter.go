package api

import (
	"go_vue/api/advert_api"
	"go_vue/api/images_api"
	"go_vue/api/menu_api"
	"go_vue/api/settings_api"
	"go_vue/api/tag_api"
	"go_vue/api/user_api"
)

type AipGroup struct {
	SettingsApi settings_api.SettingsApi //系统
	ImagesApi   images_api.ImagesApi     //图片
	AdvertApi   advert_api.AdvertApi     //广告
	MenuApi     menu_api.MenuApi         //菜单
	UserApi     user_api.UserApi         //用户
	TagApi      tag_api.TagApi           //标签
}

var ApiGroupApp = new(AipGroup)
