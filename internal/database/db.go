package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type UserInformation struct {
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

func loadCredentials() LOGIN {
	err := godotenv.Load("../secret.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return LOGIN{
		User: os.Getenv("DATABASE_USER"),
		Pass: os.Getenv("DATABASE_PASSWORD"),
		Ip:   os.Getenv("DATABASE_IP"),
		Port: os.Getenv("DATABASE_PORT"),
		Name: os.Getenv("DATABASE_NAME"),
	}
}

func ConnectDB() (*sql.DB, error) {
	credentials := loadCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.User, credentials.Pass, credentials.Ip, credentials.Port, credentials.Name)

	fmt.Println(credentials.Ip, credentials.User, credentials.Ip, credentials.Port, credentials.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database")
	}

	// Check connection
	if err = db.Ping(); err != nil {
		fmt.Println("failed to ping the server")
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connected to MySQL")
	return db, nil
}

func GetUser(userName string) (reg.UserInformation, error) {
	var userInfo reg.UserInformation
	query := "SELECT user.id FROM user WHERE name LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return reg.UserInformation{}, fmt.Errorf("failed to connect to the database")
	}

	row := conn.QueryRow(query, userName)
	err = row.Scan(&userInfo.UserName)
	if err != nil {
		return reg.UserInformation{}, fmt.Errorf("there was an error in retrieving the information from the database")
	}
	return userInfo, nil
}

func CheckPass(pass string, data reg.UserInformation) bool {
	query := "SELECT ? FROM ? WHERE ? LIKE ?"

	conn, err := db.ConnectDB()
	if err != nil {
		return false
	}

	queryError := conn.QueryRow(query, data.UserID, "USER", data.UserPassword, pass)
	return queryError == nil
}

// Stats returns the database statistics
func Stats(db *sql.DB) sql.DBStats {
	return db.Stats()
}

func CloseDb(db *sql.DB) {
	db.Close()
}
