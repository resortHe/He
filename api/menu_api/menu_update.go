package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	//先把之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		//操作第三章表 创建banner
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err := global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("创建菜单图片失败", c)
			return
		}
	}
	//普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OkWithMsg("修改菜单成功", c)
}
