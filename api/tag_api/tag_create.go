package tag_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` //显示标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复判断
	var advert models.TagModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该标签已存在", c)
		return
	}
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加标签失败", c)
		return
	}
	res.OkWithMsg("添加标签成功", c)
}
