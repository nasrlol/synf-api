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
	w.Header().Set("Content-type", "application/json")

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := `SELECT * FROM USER WHERE name LIKE ? AND	email LIKE ? AND password LIKE ?`

	conn, _ := database.Connect()

	var user models.UserInformation
	var hashedPassword []byte

	err = conn.QueryRow(query, request.Email, hashedPassword).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Verified)

	switch {
	case err == sql.ErrNoRows:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case err != nil:

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password))
	if err != nil {
		return
	}

	user.Password = ""

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
}
