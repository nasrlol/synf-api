package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "synf/database"

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

type userInformation struct {
	UserID            uint8  `json:"id"`
	UserName          string `json:"user_name"`
	UserRole          bool   `json:"user_role"`
	UserEmail         string `json:"user_email"`
	UserEmailVerified bool   `json:"is_email_verified"`
	UserPassword      string `json:"user_password"`
}

type LOGIN struct {
	User string
	Pass string
	Ip   string
	Port string
	Name string
}

func UserReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user userInformation

	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("converting user information to json...")
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	fmt.Println("hasing the user password...")
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

func insertUser(data userInformation) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	if db != nil {
		defer db.Close()
	} else {
		return fmt.Errorf("db is nil 501")
	}
	

	query := `INSERT INTO USER (user_name, user_password, user_role, user_email) VALUES (?, ?, ?, ?)`
	fmt.Println("inserting the information into the database")
	db.Exec(query, data.UserName, data.UserPassword, boolToInt(data.UserRole), data.UserEmail)

	fmt.Println("User inserted successfully!")
	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var data userInformation
	user, err := queryUser(data) 
	if err != nil {
		fmt.Printf("error") 
	}
	fmt.Println(user)
	
}

func queryUser(data userInformation) (userInformation, error) {

	db, err := db.ConnectDB()
	if err != nil{
		return userInformation{}, fmt.Errorf("database connection error when trying to pull the user information form the database")
	}

	query := "SELECT EXISTS(SELECT 1 FROM USER WHERE user_name = ?)"

	row, err := db.Query(query, data.UserName)
	if err != nil {
		return userInformation{}, fmt.Errorf("internal databse error couldn't complete query")
	}
	defer row.Close()

	var user userInformation 

	for row.Next(){
		err := row.Scan(&data.UserName, &data.UserPassword, &data.UserRole, &data.UserEmail)
		if err != nil {
			return userInformation{}, fmt.Errorf("error scanning rows") 
		}
	}

	fmt.Println("Retrieved username succesfully")
	return user, nil 
}
