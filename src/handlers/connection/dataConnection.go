package dataConnection

import (
	"log"
	"net/http"
	"time"
	sys "synf/handlers/data"
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
	defer conn.Close()

	for data := range sys.CpuTemperature() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(data)); err != nil {
			log.Printf("error writing the message over the websocket: %v\n", err)
			break
		}
	}
}

func Init() {
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
