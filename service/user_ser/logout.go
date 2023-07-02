package user_ser

import (
	"go_vue/service/redis_ser"
	"go_vue/utils/jwts"
	"time"
)

// Logout 将token和过期时间传递到redis
func (UserSerVice) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)

	return redis_ser.Logout(token, diff)
}
