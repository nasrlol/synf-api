package routes

import (
	"net/http"
)

func init() {
	http.HandleFunc("/api/data/cpu", CpuHandler)
	http.HandleFunc("/api/data/cpu", RamHandler)
	http.HandleFunc("/api/data/disk", DiskHandler)
	http.HandleFunc("/api/data/ram", DeviceHandler)
	http.HandleFunc("/api/about", AboutHandler)
}

func CpuHandler(w http.ResponseWriter, r *http.Request) {
}

func RamHandler(w http.ResponseWriter, r *http.Request) {
}

func DiskHandler(w http.ResponseWriter, r *http.Request) {
}

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {}
