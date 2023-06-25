package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
