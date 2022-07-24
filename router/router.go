package router

import (
	"code.project.com/InstantMessaging/router/middleware"
	"code.project.com/InstantMessaging/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户登录
	r.POST("/login", service.UserLogin)
	// 用户信息
	rg := r.Group("/user", middleware.AuthLoginCheck())
	{
		rg.GET("/detail", service.UserDetail)
	}
	return r
}
