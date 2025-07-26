package routes

import (
	"net/http"

	h "synf/internal/api/data/rest"

	"github.com/gorilla/mux"
)

func InitRestRoutes() {
	r := mux.NewRouter()

	/*
		r.HandleFunc("/data/ram", h.CreateRam).Methods("POST")
		r.HandleFunc("/data/ram/{id}", h.GetRam).Methods("GET")
		r.HandleFunc("/data/ram/{id}", h.UpdateRam).Methods("PUT")
		r.HandleFunc("/data/ram/{id}", h.DeleteRam).Methods("DELETE")

		r.HandleFunc("/data/disk", h.CreateDisk).Methods("POST")
		r.HandleFunc("/data/disk", h.GetDisk).Methods("GET")
		r.HandleFunc("/data/disk", h.UpdateDisk).Methods("PUT")
		r.HandleFunc("/data/disk", h.DeleteDisk).Methods("DELETE")

		r.HandleFunc("/data/device", h.CreateDevice).Methods("POST")
		r.HandleFunc("/data/device", h.GetDevice).Methods("GET")
		r.HandleFunc("/data/device", h.UpdateDevice).Methods("PUT")
		r.HandleFunc("/data/device", h.DeleteDevice).Methods("DELETE")
	*/

	r.HandleFunc("/user/login", h.GetUser).Methods("POST")
	r.HandleFunc("/user/registration", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("POST")

	http.Handle("/", r)
}
