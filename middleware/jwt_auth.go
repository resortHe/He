package middleware

import (
	"github.com/gin-gonic/gin"
	"go_vue/models/ctype"
	"go_vue/models/res"
	"go_vue/service/redis_ser"
	"go_vue/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带Token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}

		//登录的用户
		c.Set("claims", claims)
	}
}

// JwtAdmin 管理才可以调用的接口中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带Token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			res.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}

		//登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
