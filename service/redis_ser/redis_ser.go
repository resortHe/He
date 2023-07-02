package redis_ser

import (
	"context"
	"go_vue/global"
	"go_vue/utils"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(context.Background(), prefix+token, "", diff).Err()
	return err
}
func CheckLogout(token string) bool {
	val := global.Redis.Keys(context.Background(), prefix+"*").Val()
	if utils.InList("logout_"+token, val) {
		return true
	}
	return false
}
