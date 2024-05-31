package ping

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func List(ctx *gin.Context) {
	// ping 接口测试
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "time": nowTime, "msg": "pong"})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ping(ctx *gin.Context) {
	// 升级req to websocket
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, gin.H{"msg": "xxxx"})
		return
	}
	defer ws.Close()
	for {
		// msgType 1 文本消息 2二进制 8 关闭消息
		msgType, msgContent, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err, "接受消息失败")
			continue
		}
		fmt.Println(string(msgContent), "消息内容")
		if msgType == websocket.CloseMessage {
			break
		}
		// 创建消息
		msgResponse := fmt.Sprintf("%s 收到了!", string(msgContent))
		err = ws.WriteMessage(msgType, []byte(msgResponse))
		if err != nil {
			fmt.Println(err, "发送消息失败")
		}

	}
}
