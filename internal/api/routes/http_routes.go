package routes

import (
	"net/http"
	handlers "synf/internal/api/data/rest"
)

func InitRestRoutes() {
	http.HandleFunc("/data/ram", ramHandler)
	http.HandleFunc("/data/disk", diskHandler)
	http.HandleFunc("/data/device", deviceHandler)
	http.HandleFunc("/user/login", handlers.GetUser)
	http.HandleFunc("/user/registration", handlers.CreateUser)
}
