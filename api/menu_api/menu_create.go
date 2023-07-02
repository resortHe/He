package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/ctype"
	"go_vue/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                //简介的切换的时间 单位s
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    //Banner的切换的时间 单位s
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          //具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ? ", cr.Title, cr.Path).RowsAffected
	fmt.Println(count)
	if count > 0 {
		res.FailWithMsg("重复的菜单", c)
		return
	}

	//创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMsg("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel

	for _, sort := range cr.ImageSortList {
		//这里也得判断imageId是否真的有图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	//给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单图片关联失败失败", c)
		return
	}
	res.OkWithMsg("菜单添加成功", c)
}
