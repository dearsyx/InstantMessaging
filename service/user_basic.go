package service

import (
	"code.project.com/InstantMessaging/models"
	"code.project.com/InstantMessaging/pkg/token"
	"code.project.com/InstantMessaging/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	tokenString, err := token.GenerateToken(user.Identity, user.Username, user.Email)
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

// UserDetail 获取用户详细信息
func UserDetail(c *gin.Context) {
	userClaim, exist := c.Get("user_claim")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
	}
	user := userClaim.(*token.UserClaims)
	userBasic, err := models.GetUserByIdentity(user.UserID)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查找用户失败",
		})
		return
	}
	// 找到了用户
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  fmt.Sprintf("用户 %s 登录成功", userBasic.Username),
		"data": userBasic,
	})
}
