package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"go_vue/global"
	"time"
)

// GenToken 获取token
func GenToken(user JwtPayload) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claims := CustomClaims{
		JwtPayload: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(MySecret)
}
