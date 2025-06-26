package Handlers 

import (
	"fmt"
	"log"
	"net/http"
	"synf/internal/api/data"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		//		return r.Host == "nsrddyn.com"
		return true
	},
}

func wsHandlerCPU(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket :%v\n", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	for dt := range data.CpuTemperature() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(dt)); err != nil {
			log.Printf("error writing the message over the websocket: %v\n", err)
			break
		}
	}
}

func restHandlerGPU(w http.ResponseWriter, r *http.Request) {}

func InitCPU() {
	fmt.Println("initalizing the websocket and serving the cpu handler")
	server := &http.Server{
		Addr:           ":8085",
		Handler:        http.HandlerFunc(wsHandlerCPU),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to listen and serve the server %v\n", err)
	}
}
