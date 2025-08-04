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

/*
	body, _ := io.ReadAll(r.Body)
	fmt.Println("RAW BODY RECEIVED:", string(body))
	r.Body = io.NopCloser(bytes.NewReader(body)) // reset for decode
*/

func GetUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `SELECT * FROM users WHERE email = ?`

	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

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

	database.Close(conn)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Verfied  int    `jons:"verified"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
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

	query := `INSERT INTO users (name, email, password, role, verified) VALUES (?, ?, ?, ?, ?)`

	_, err = conn.Exec(query, request.Name, request.Email, hashedPassword, request.Role, request.Verfied)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusOK)
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

	query := `DELETE FROM users WHERE id = ?;`
	conn, err := database.Connect()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		return
	}
	defer database.Close(conn)
	_, err = conn.Exec(query, id)
	if err != nil {
		http.Error(w, "Unable to execute the query", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
