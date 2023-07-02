package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models/ctype"
	"go_vue/models/res"
	"go_vue/service/user_ser"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  //昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` //用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   //密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请输入权限"`       //权限
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err = user_ser.UserSerVice{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("用户%s创建成功", cr.UserName), c)
}
