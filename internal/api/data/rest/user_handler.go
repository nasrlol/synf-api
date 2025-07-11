package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"synf/internal/api/data/models"

	"synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
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

	var user models.UserInformation

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
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
