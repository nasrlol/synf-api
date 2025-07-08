package routes

import (
	"log"
	"net/http"
	"time"
)

func InitWsRoutes() {
	// http.HandleFunc("cpu", ws.MakeWsHandler(services.Cpu()))

	server := &http.Server{
		Addr:           ":8085",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start WebSocket server: %v\n", err)
	}
}
