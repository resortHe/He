package user_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/models"
	"go_vue/models/ctype"
	"go_vue/models/res"
	"go_vue/service/common"
	"go_vue/utils/desens"
	"go_vue/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//管理员才可见用户名称
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OkWithList(users, count, c)
}
