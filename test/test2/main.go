package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		fmt.Println("pong")
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)
		conn.WriteMessage(websocket.PingMessage, nil)
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	//go func() {
	//	for {
	//		connections.Range(func(key, value interface{}) bool {
	//			conn := key.(*websocket.Conn)
	//			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
	//				fmt.Println("Error sending ping:", err)
	//				connections.Delete(key)
	//			}
	//			return true
	//		})
	//		time.Sleep(10 * time.Second)
	//	}
	//}()

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
