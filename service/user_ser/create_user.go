package user_ser

import (
	"errors"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/ctype"
	"go_vue/utils/pwd"
)

const Avatar = "/uploads/avatar/default.png"

func (UserSerVice) CreateUser(userName string, nickName string, password string, role ctype.Role, email string, ip string) error {
	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ? ", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在,请重新输入")
		return errors.New("用户名已存在")
	}
	//对密码进行hash

	hashPwd := pwd.HashPwd(password)
	//头像
	//1默认头像
	//2随机头像
	//入库
	err = global.DB.Create(&models.UserModel{
		NickName:  nickName,
		UserName:  userName,
		Password:  hashPwd,
		Email:     email,
		Role:      role,
		Avatar:    Avatar,
		IP:        ip,
		Addr:      "内网地址",
		SigStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
