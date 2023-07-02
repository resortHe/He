package service

import (
	"go_vue/service/image_ser"
	"go_vue/service/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserSerVice  user_ser.UserSerVice
}

var ServiceApp = new(ServiceGroup)
