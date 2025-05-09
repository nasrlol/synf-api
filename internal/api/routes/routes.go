package routes

import (
	"net/http"
)

func init() {
	http.HandleFunc("/cpu", CpuHandler)
	http.HandleFunc("/gpu", GpuHandler)
	http.HandleFunc("/ram", RamHandler)
	http.HandleFunc("/disk", DiskHandler)
	http.HandleFunc("/gendev", GenDevHandler)
}

func CpuHandler(w http.ResponseWriter, r *http.Request) {

}

func GpuHandler(w http.ResponseWriter, r *http.Request) {

}

func RamHandler(w http.ResponseWriter, r *http.Request) {

}

func DiskHandler(w http.ResponseWriter, r *http.Request) {

}

func GenDevHandler(w http.ResponseWriter, r *http.Request) {

}
