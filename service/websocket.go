package service

import (
	"code.project.com/InstantMessaging/models"
	"code.project.com/InstantMessaging/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var wsConn = make(map[string]*websocket.Conn)

// WebsocketMessage 发送和接收消息
func WebsocketMessage(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "建立连接失败" + err.Error(),
		})
		return
	}
	defer func() { _ = conn.Close() }()
	uc, exist := ctx.Get("user_claim")
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户未登录",
		})
		return
	}
	userClaim := uc.(*token.UserClaims)
	wsConn[userClaim.UserID] = conn
	for {
		ms := new(models.MessageStruct)
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("read error : %s \n", err.Error())
			return
		}
		// 判断用户是否属于消息体的房间
		_, err = models.GetUserRoomByUserAndRoomIdentity(userClaim.UserID, ms.RoomIdentity)
		if err != nil {
			log.Printf("用户 %s 不在房间 %s 内\n", userClaim.UserID, ms.RoomIdentity)
			return
		}
		// 保存消息
		_ = models.SaveMessageToMongo(userClaim.UserID, ms)
		// 获取特定房间的在线用户
		userGroup, err := models.GetUserGroupByRoomID(ms.RoomIdentity)
		if err != nil {
			log.Printf("[DB ERROR] : %s \n", err.Error())
			return
		}
		for _, user := range userGroup {
			if conn, ok := wsConn[user.UserIdentity]; ok {
				_ = conn.WriteMessage(websocket.TextMessage, []byte(ms.Message))
			}
		}
	}
}
