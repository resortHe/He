package advert_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        //显示标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     //跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` //图片
	IsShow bool   `json:"is_show"   structs:"is_show"`                                 //是否显示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description  创建广告
// @Param data body  AdvertRequest      true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该广告已存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加广告失败", c)
		return
	}
	res.OkWithMsg("添加广告成功", c)
}
