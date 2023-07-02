package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"` //主键ID
	CreatedAt time.Time ` json:"created_at"`          //创建时间
	UpdatedAt time.Time `json:"-"`                    //更新时间
}
type PageInfo struct {
	Page  int    `form:"page"`  //页数
	Key   string `form:"key"`   //模糊查询的关键字
	Limit int    `form:"limit"` //每页显示的数量
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
