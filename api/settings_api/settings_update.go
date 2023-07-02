package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/config"
	"go_vue/global"
	"go_vue/models/res"
)

// SettingsInfoUpdateView 优点减少了接口数量，缺点入参出参不统一
// @Tags 系统管理
// @Summary 更新配置信息
// @Description  根据名称更新对应的配置信息
// @Param name path string true "配置项名称" Enums(site, email, qq, qiniu, jwt)
// @Router /api/settings/{name} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = email
	case "qq":
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = qq
	case "qiniu":
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = qiniu
	case "jwt":
		var jwt config.Jwt
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = jwt
	default:
		res.FailWithMsg("没有对于的配置信息", c)

	}
}
