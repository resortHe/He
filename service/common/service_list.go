package common

import (
	"go_vue/global"
	"go_vue/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// ComList 分页查询
func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按时间往前排
	}
	query := DB.Where(model)
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model) //这里query会受上面的查询影响 这是复位操作
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
