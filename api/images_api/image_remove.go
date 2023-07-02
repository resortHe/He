package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

// ImageRemoveView 图片删除
// @Tags 图片管理
// @Summary 批量删除图片
// @Description  批量删除图片
// @Param data body  models.RemoveRequest     true "图片id列表"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel
	affected := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if affected == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMsg(fmt.Sprintf("共删除%d张图片", affected), c)
}
