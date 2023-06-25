package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"ID"` //主键ID
	CreatedAt time.Time `json:"Created_at"`           //创建时间
	UpdateAt  time.Time `json:"-"`                    //更新时间
}
