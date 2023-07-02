package user_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
	"go_vue/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` //旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     //新密码

}

// UserUpdatePasswordView 修改登录人的ID
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//判断密码是否一致
	if !pwd.CheckPwd(userModel.Password, cr.OldPwd) {
		res.FailWithMsg("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&userModel).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("密码修改错误", c)
		return
	}
	res.OkWithMsg("密码修改成功", c)
	return
}
