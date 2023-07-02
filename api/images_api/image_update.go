package images_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// ImageUpdateView   更新图片
// @Tags 图片管理
// @Summary 更新图片
// @Description  更新图片
// @Param data body  ImageUpdateRequest    true "图片的名字"
// @Router /api/images/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithMsg("名称修改成功", c)
	return
}
