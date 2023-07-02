package user_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
	"go_vue/utils/pwd"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 用户名或邮箱登录
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//校验登录
	checkPwd := pwd.CheckPwd(userModel.Password, cr.Password)
	if !checkPwd {
		global.Log.Warn("用户名密码错误")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//登录成功 生成token
	token, err := jwts.GenToken(jwts.JwtPayload{
		Nickname: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	res.OkWithData(token, c)

}
