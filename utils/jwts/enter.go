package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayload jwt中payload数据
type JwtPayload struct {
	//Username string `json:"username"` //用户名
	Nickname string `json:"nickname"` //昵称
	Role     int    `json:"role"`     //权限 1 管理员,2 普通用户,3 游客
	UserID   uint   `json:"user_id"`  //用户ID
}

var MySecret []byte

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}
