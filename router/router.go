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
	// 发送验证码
	r.POST("/captcha", service.SendCode)
	// 用户信息
	rg := r.Group("/user", middleware.AuthLoginCheck())
	{
		// 用户详细信息
		rg.GET("/detail", service.UserDetail)
		// 发送和接收消息
		rg.GET("/message", service.WebsocketMessage)
	}
	return r
}
