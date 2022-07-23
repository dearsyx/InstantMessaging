package service

import (
	"code.project.com/InstantMessaging/models"
	"code.project.com/InstantMessaging/pkg/token"
	"code.project.com/InstantMessaging/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	// 如果用户输入为空
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名和密码不能为空",
		})
		return
	}
	// 根据账户密码取数据
	user, err := models.GetUserByAccountPassword(account, util.GenerateMD5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	// 生成token
	tokenString, err := token.GenerateToken(user.Account, user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "状态保存失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": tokenString,
		},
	})
	return
}
