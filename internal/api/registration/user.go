package registration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	db "synf/internal/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

type UserInformation struct {
	UserID            uint8  `json:"id"`
	UserName          string `json:"name"`
	UserEmail         string `json:"email"`
	UserPassword      string `json:"password"`
	UserRole          bool   `json:"role"`
	UserEmailVerified bool   `json:"verified"`
	UserIsLogging     bool   `json:"is_logging"`
}

type LOGIN struct {
	User string
	Pass string
	Ip   string
	Port string
	Name string
}

func insertUser(data UserInformation) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	if conn != nil {
		defer func(conn *sql.DB) {
			err := conn.Close()
			if err != nil {
			}
		}(conn)
	} else {
		return fmt.Errorf("db is nil 501")
	}

	query := `INSERT INTO USER (name, password, user_role, email) VALUES (?, ?, ?, ?)`
	fmt.Println("inserting the information into the database")
	conn.Exec(query, data.UserName, data.UserPassword, boolToInt(data.UserRole), data.UserEmail)

	fmt.Println("User inserted successfully!")
	return nil
}

func UserReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user UserInformation

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

	if err := insertUser(user); err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
	})
}

func selectUser(data UserInformation) UserInformation {
	db, err := db.ConnectDB()
	if err != nil {
		return UserInformation{}
	}

	query := "SELECT EXISTS(SELECT 1 FROM USER WHERE id = ?)"
	row := db.QueryRow(query, string(data.UserID))

	err = row.Scan(&data.UserName, &data.UserPassword, &data.UserRole, &data.UserEmail)
	if err != nil {
		return UserInformation{}
	} else {
		return UserInformation{}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data UserInformation
	data = selectUser(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
