package message_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
)

func (MessageApi) MessageRemoveView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var messageList = make([]models.MessageModel, 0)
	var _messageList []models.MessageModel
	var affected int64
	if claims.UserID == cr.UserID {
		affected = global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? and rev_user_id = ? ", claims.UserID, claims.UserID).RowsAffected
	} else {
		affected = global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? or rev_user_id = ? ", claims.UserID, claims.UserID).RowsAffected
	}
	if affected == 0 {
		res.FailWithMsg("聊天记录为空", c)
		return
	}
	for _, model := range _messageList {
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	global.DB.Model(&messageList).Delete(&messageList)
	res.OkWithMsg("清空聊天记录成功", c)
	return
}
