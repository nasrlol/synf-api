package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

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

func loadCredentials() LOGIN {
	err := godotenv.Load("secret.env")
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

func connectDB() (*sql.DB, error) {
	credentials := loadCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.User, credentials.Pass, credentials.Ip, credentials.Port, credentials.Name)

	fmt.Println(credentials.Ip, credentials.User, credentials.User)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
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

	fmt.Println("Connected to MySQL")
	return db, nil
}

func userReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO USER (user_name, user_password, user_role, user_email) VALUES (?, ?, ?, ?)`
	fmt.Println("inserting the information into the database")
	db.Exec(query, data.UserName, data.UserPassword, boolToInt(data.UserRole), data.UserEmail)

	fmt.Println("User inserted successfully!")
	return nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("Server is running"))
}

func serveServer() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/user/register", userReg)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getOutboundIp() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func raw_connect(host string,  port string) {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
}

func main() {
	fmt.Println("API STARTED...")
	fmt.Println(getOutboundIp())
	raw_connect("localhost", "8080")
	serveServer()
}
