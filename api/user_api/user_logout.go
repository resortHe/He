package user_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models/res"
	"go_vue/service"
	"go_vue/utils/jwts"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	//需要计算距离现在的过期时间
	err := service.ServiceApp.UserSerVice.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OkWithMsg("注销成功", c)
}
