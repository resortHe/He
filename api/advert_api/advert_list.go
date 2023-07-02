package advert_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/service/common"
	"strings"
)

// AdvertListView  广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description  广告列表
// @Param data query  models.PageInfo      false "查询参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	//判读Referer是否包含admin，如果是全部返回，不是就返回is_show等于true
	header := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(header, "admin") {
		//admin来的
		isShow = false
	}
	list, count, err := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
