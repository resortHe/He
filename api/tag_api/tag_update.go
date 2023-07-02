package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	var id = c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMsg("该标签不存在", c)
		return
	}
	//structs结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.OkWithMsg("修改标签成功", c)
}
