package models

// LoginDateModel 统计用户登录的信息
type LoginDateModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` //登录的ip
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登录设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
