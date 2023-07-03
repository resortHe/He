package message_api

import (
	"github.com/gin-gonic/gin"
	"go_vue/global"
	"go_vue/models"
	"go_vue/models/res"
	"go_vue/utils/jwts"
	"time"
)

type Message struct {
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"` //发送人id
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `gorm:"primaryKey" json:"rev_user_id"` //接收人id
	RevUserNickName  string    `gorm:"size:42" json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`    //消息内容
	CreatedAt        time.Time `json:"created_at"` //创建时间
	MessageCount     int       `json:"message_count"`
}
type MessageGroup map[uint]*Message

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var messageList []models.MessageModel
	var messageGroup = MessageGroup{}
	var messages []Message
	global.DB.Order("created_at asc").Find(&messageList, "send_user_id = ? or rev_user_id = ? ", claims.UserID, claims.UserID)
	for _, model := range messageList {
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			//不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}
	res.OkWithData(messages, c)

}
