package routes

import (
	"net/http"

	handlers "synf/internal/api/data/rest"
)

func InitRestRoutes() {
	http.HandleFunc("/data/ram", RamHandler)
	http.HandleFunc("/data/disk", DiskHandler)
	http.HandleFunc("/data/device", DeviceHandler)
	http.HandleFunc("/user/login", handlers.LoginHandler)
	http.HandleFunc("/user/registration", handlers.CreateUser)
}
