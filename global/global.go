package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go_vue/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Log      *logrus.Logger
	Config   *config.Config
	DB       *gorm.DB
	MysqlLog logger.Interface
	Redis    *redis.Client
)
