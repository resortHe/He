package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

// AdvertRemoveView  批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description  批量删除广告
// @Param data body  models.RemoveRequest     true "广告id列表"
// @Router /api/advert [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var advertList []models.AdvertModel
	affected := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if affected == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMsg(fmt.Sprintf("共删除%d个广告", affected), c)
}
