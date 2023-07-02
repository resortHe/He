package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var userList []models.UserModel
	affected := global.DB.Find(&userList, cr.IDList).RowsAffected
	if affected == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//删除第三张表
	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除用户，消息表，评论表，用户收藏文章，用户发布的文章都要删
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除用户失败", c)
		return
	}

	res.OkWithMsg(fmt.Sprintf("共删除%d个用户", affected), c)
}
