package core

import (
	"fmt"
	"go_vue/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s]mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //最大复用时间不能超过mysql的wait_time
	return db
}