package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"synf/internal/api/data/models"

	db "synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func insertUser(data models.UserInformation) (bool, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to the database")
	}
	if conn != nil {
		defer func(conn *sql.DB) {
			err := conn.Close()
			if err != nil {
			}
		}(conn)
	} else {
		return false, fmt.Errorf("db is nil")
	}

	query := `INSERT INTO USER (name, password, user_role, email) VALUES (?, ?, ?, ?)`
	fmt.Println("inserting the information into the database")
	_, err = conn.Exec(query, data.UserName, data.UserPassword, data.UserRole, data.UserEmail)
	if err != nil {
		return false, fmt.Errorf("failed to execute the query")
	}
	return true, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserInformation

	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("converting user information to json...")
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	fmt.Println("hashing the user password...")
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.UserPassword = string(hashedPassword)

	if value, err := insertUser(user); value != true && err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	request := r.Body(models.UserInformation{
		UserName:  "",
		UserEmail: "",
		UserRole:  "",
	})

	w.Header().Set("Content-Type", "application/json")

	conn, err := db.ConnectDB()
	if err != nil {
		w.WriteHeader(404)
	}
	query := `SELECT EXISTS (
					SELECT 1
					FROM USER
					WHERE name LIKE ?
					AND
					email LIKE ?
					AND
					password LIKE ?)`

	conn.Exec(query, request.Username, request)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
