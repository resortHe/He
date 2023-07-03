package message_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/service/common"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	list, count, err := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	//需要展示标签下的文章数目
	res.OkWithList(list, count, c)
}
