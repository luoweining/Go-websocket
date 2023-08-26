package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		// 升级HTTP连接为WebSocket连接
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade WebSocket connection:", err)
			return
		}
		defer conn.Close()

		for {
			// 读取客户端发送的消息
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message:", err)
				break
			}

			log.Printf("Received message: %s\n", msg)

			// 发送消息给客户端
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Failed to write message:", err)
				break
			}
		}
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
