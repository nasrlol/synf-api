package routes

import (
	"log"
	"net/http"
	"time"

	"synf/internal/api/data/ws/"
)

func InitWsRoutes() {
	http.HandleFunc("cpu/temperature", MakeWsHandler(data.CpuTemperature()))
	http.HandleFunc("cpu/frequency", MakeWsHandler(data.CpuFrequency()))

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
