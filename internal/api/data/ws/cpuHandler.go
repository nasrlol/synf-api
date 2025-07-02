package http

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	ser "synf/internal/api/data/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return r.Host == "api.nsrddyn.com"
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

		for dt, _ := range ser.Cpu("./cpu", "temperature") {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(dt)); err != nil {
				log.Printf("Error writing message over WebSocket: %v\n", err)
				break
			}
		}
	}
}
