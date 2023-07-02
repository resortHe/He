package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menutList []models.MenuModel
	affected := global.DB.Find(&menutList, cr.IDList).RowsAffected
	if affected == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	//删除第三张表
	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menutList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&menutList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除菜单失败", c)
		return
	}

	res.OkWithMsg(fmt.Sprintf("共删除%d个菜单", affected), c)
}
