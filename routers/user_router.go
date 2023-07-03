package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go_vue/api"
	"go_vue/middleware"
)

var store = cookie.NewStore([]byte("nynu0834"))

// UserRouter  图片路由
func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.GET("list", userApi.UserListView)
	router.POST("email_login", userApi.EmailLoginView)
	router.POST("register", userApi.UserRegisterView)
	router.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	userRouterAdmin := router.Group("user_admin")
	{
		userRouterAdmin.Use(middleware.JwtAdmin())
		userRouterAdmin.PUT("user_role", userApi.UserUpdateRoleView)
		userRouterAdmin.DELETE("delete", userApi.UserRemoveView)
		userRouterAdmin.GET("list", userApi.UserListView)
		userRouterAdmin.POST("create_user", userApi.UserCreateView)
	}
	userRouterAuth := router.Group("user_auth")
	{
		userRouterAuth.Use(middleware.JwtAuth())

		userRouterAuth.PUT("update_password", userApi.UserUpdatePasswordView)
		userRouterAuth.POST("logout", userApi.LogoutView)
		userRouterAuth.POST("email_bind", userApi.UserBindEmailView)
	}

}
