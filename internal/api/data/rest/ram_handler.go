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

func GetRam(w http.ResponseWriter, r *http.Request) {
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
		query := `SELECT * FROM rams WHERE id = ?`
		var ram models.Ram

		conn.QueryRow(query, request.Id).Scan(&ram.Id)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(ram)

	}
}

func CreateRam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == " " {
		http.Error(w, "No id", http.StatusBadRequest)
		return
	}

	var request struct {
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

	base := `UPDATE rams SET `
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

func UpdateRam(w http.ResponseWriter, r *http.Request) {
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
		query := `SELECT * FROM rams WHERE id = ?`
		var ram models.Ram

		conn.QueryRow(query, request.Id).Scan(&ram.Id)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(ram)

	}
}

func DeleteRam(w http.ResponseWriter, r *http.Request) {
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
		query := `DELETE FROM rams WHERE id = ?`

		_, err := conn.Exec(query, request.Id)
		if err != nil {
			http.Error(w, "bad request", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
