package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"synf/internal/api/data/models"

	"synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := `SELECT * FROM USER WHERE email = ?`

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	defer database.Close(conn)

	var user models.User

	err = conn.QueryRow(query, request.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Verified)

	switch {
	case err == sql.ErrNoRows:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case err != nil:

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	user.Password = ""

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	defer database.Close(conn)

	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	_, err = conn.Exec(query, request.Name, request.Email, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	database.Close(conn)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == " " {
		http.Error(w, "No id", http.StatusBadRequest)
		return
	}

	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Verified bool   `json:"verified"`
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

	base := `UPDATE users SET `
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == " " {
		http.Error(w, "No id", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM users WHERE ?;`
	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}
	defer database.Close(conn)
	conn.Exec(query, id)
}
