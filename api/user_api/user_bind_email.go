package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/plugins/email"
	"go_vue/utils/jwts"
	"go_vue/utils/pwd"
	"go_vue/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password" `
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	//用户绑定邮箱，第一次输入是邮箱
	//后台给邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次发验证码
		//生产4位验证码，将生产的验证码存入session
		code := random.Code(4)
		fmt.Println(code)
		//写入session
		session.Set("valid_code", code)
		session.Set("email", cr.Email)
		err := session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是"+code)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("验证码发送失败", c)
			return
		}
		res.OkWithMsg("验证码已发送", c)
		return
	}
	//第二次 用户输入邮箱验证码 密码
	code := session.Get("valid_code")
	Email := session.Get("email")
	//校验验证码
	if code != *cr.Code {
		res.FailWithMsg("验证码错误", c)
		return
	}
	//修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMsg("密码长度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	//第一次的邮箱和第二次的邮箱也要做一致性校验
	if Email != cr.Email {
		res.FailWithMsg("请输入对应邮箱的验证码", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		res.FailWithMsg("绑定邮箱失败", c)
		return
	}
	//完成绑定
	res.OkWithMsg("邮箱绑定成功", c)
	return

}
