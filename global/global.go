package global

import (
	"github.com/sirupsen/logrus"
	"go_vue/config"
	"gorm.io/gorm"
)

var (
	Log    *logrus.Logger
	Config *config.Config
	DB     *gorm.DB
)
