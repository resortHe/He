package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

// AdvertUpdateView  更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description  更新广告
// @Param data body  AdvertRequest    true "广告的一些参数"
// @Router /api/advert/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var id = c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMsg("该广告不存在", c)
		return
	}
	//structs结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OkWithMsg("修改广告成功", c)
}
