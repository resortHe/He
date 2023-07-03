package message_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` //发送人id
	RevUserId  uint   `json:"rev_user_id" binding:"required"`  //接收人ID
	Context    string `json:"context" binding:"required"`      //内容
}

func (MessageApi) MessageCreateView(c *gin.Context) {
	//SendUserID 就是当前登入人的id
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if claims.UserID != cr.SendUserID {
		res.FailWithMsg("请先登录再发送信息", c)
		return
	}
	var sendUser, revUse models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMsg("发送者ID不存在", c)
		return
	}
	err = global.DB.Take(&revUse, cr.RevUserId).Error
	if err != nil {
		res.FailWithMsg("接收者ID不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserId,
		RevUserNickName:  revUse.NickName,
		RevUserAvatar:    revUse.Avatar,
		IsRead:           false,
		Content:          cr.Context,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("消息发送失败", c)
		return
	}
	res.OkWithMsg("消息发送成功", c)
	return
}
