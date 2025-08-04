package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"synf/internal/api/data/models"
	"synf/internal/api/data/services"
	"synf/internal/database"

	"github.com/gorilla/mux"
)

func GetDisk(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Size  int    `json:"size"`
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	defer database.Close(conn)

	if services.Check_session_valid(request.Token, conn) {
		query := `SELECT * FROM disks WHERE id = ?`
		var device models.Device

		conn.QueryRow(query, request.Id).Scan(&device.Id, &device.Name)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(device)

	}
}

func CreateDisk(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Size  int    `json:"size"`
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Failed hashing the password", http.StatusInternalServerError)
		return
	}

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	defer database.Close(conn)

	query := `INSERT INTO disks (name, size) VALUES (?, ?))`

	_, err = conn.Exec(query, request.Name, request.Size)
	if err != nil {

		fmt.Println("failed here")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
}

func UpdateDisk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == " " {
		http.Error(w, "No id", http.StatusBadRequest)
		return
	}

	var request struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := reflect.ValueOf(request)
	typ := reflect.TypeOf(request)
	entries := make(map[string]string)

	var query strings.Builder

	base := `UPDATE devices SET `
	query.WriteString(base)

	for i := range data.NumField() {
		field := data.Field(i)
		if field.IsValid() && !field.IsZero() {
			valueStr := fmt.Sprintf("%v", field.Interface())
			entries[valueStr] = typ.Field(i).Name
		}
	}

	for key, value := range entries {
		_, _ = query.WriteString(key)
		_, _ = query.WriteString("=")
		_, _ = query.WriteString(value)
		_, _ = query.WriteString(",")
	}

	query.WriteString(`WHERE id = ?`)

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	conn.Exec(query.String(), id)
}

func DeleteDisk(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Id    int    `json:"id"`
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	defer database.Close(conn)

	if services.Check_session_valid(request.Token, conn) {
		query := `DELETE FROM devices WHERE id = ?`

		_, err := conn.Exec(query, request.Id)
		if err != nil {
			http.Error(w, "bad request", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
