package user_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/ctype"
	"go_vue/models/res"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4 " msg:"权限参数错误"`
	NickName string     `json:"nick_name"` //防止用户昵称非法,管理员有权力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户ID错误"`
}

// UserUpdateRoleView 管理员修改用户的权限
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("用户ID错误,用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改权限失败", c)
		return
	}
	res.OkWithMsg("修改权限成功", c)
}
