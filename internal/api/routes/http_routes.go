package routes

import (
	h "synf/internal/api/data/rest"

	"github.com/gorilla/mux"
)

func InitRestRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/data/ram", h.CreateRam).Methods("POST")
	r.HandleFunc("/data/ram/{id}", h.GetRam).Methods("GET")
	r.HandleFunc("/data/ram/{id}", h.UpdateRam).Methods("PATCH")
	r.HandleFunc("/data/ram/{id}", h.DeleteRam).Methods("DELETE")

	r.HandleFunc("/data/disk", h.CreateDisk).Methods("POST")
	r.HandleFunc("/data/disk", h.GetDisk).Methods("GET")
	r.HandleFunc("/data/disk", h.UpdateDisk).Methods("PATCH")
	r.HandleFunc("/data/disk", h.DeleteDisk).Methods("DELETE")

	r.HandleFunc("/data/device", h.CreateDevice).Methods("POST")
	r.HandleFunc("/data/device", h.GetDevice).Methods("GET")
	r.HandleFunc("/data/device", h.UpdateDevice).Methods("PATCH")
	r.HandleFunc("/data/device", h.DeleteDevice).Methods("DELETE")

	r.HandleFunc("/user/login", h.GetUser).Methods("POST")
	r.HandleFunc("/user/registration", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("user/{id}", h.DeleteUser).Methods("DELETE")

	return r
}
