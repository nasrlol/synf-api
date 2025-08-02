package http

import (
	"log"
	"net/http"

	"synf/internal/api/data/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MakeWsHandler(<-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v\n", err)
			return
		}
		defer func() {
			if err := conn.Close(); err != nil {
				log.Printf("Error closing WebSocket connection: %v\n", err)
			}
		}()

		for dt := range services.Cpu("./cpu", "temperature") {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(dt)); err != nil {
				log.Printf("Error writing message over WebSocket: %v\n", err)
				break
			}
		}
	}
}
