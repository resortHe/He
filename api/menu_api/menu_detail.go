package menu_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

// MenuDetailView 菜单详情
func (MenuApi) MenuDetailView(c *gin.Context) {
	//先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	//查连接表
	var menuBanners []models.MenuBannerModel
	//Preload 连表查
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id =  ?", id)
	banners := make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
	return
}
