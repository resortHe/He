package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var tagList []models.TagModel
	affected := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if affected == 0 {
		res.FailWithMsg("标签不存在", c)
		return
	}
	//如果标签下有文章，
	global.DB.Delete(&tagList)
	res.OkWithMsg(fmt.Sprintf("共删除%d个标签", affected), c)
}
