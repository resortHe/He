package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView   系统列表
// @Tags 系统管理
// @Summary 系统列表
// @Description  系统列表
// @Param name path string true "配置项名称" Enums(site, email, qq, qiniu, jwt)
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMsg("没有对于的配置信息", c)

	}
}
