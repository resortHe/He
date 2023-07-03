package message_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入用户id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	if claims.UserID == cr.UserID {
		global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? and rev_user_id = ? ", claims.UserID, claims.UserID)
	} else {
		global.DB.Order("created_at asc").Find(&_messageList, "send_user_id = ? or rev_user_id = ? ", claims.UserID, claims.UserID)
	}
	for _, model := range _messageList {
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	//点开消息，里面的每条消息从未读变成已读
	global.DB.Model(&messageList).Update("is_read", true)
	res.OkWithData(messageList, c)
	return

}
