package router

import (
	"code.project.com/InstantMessaging/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", service.UserLogin)

	return r
}
