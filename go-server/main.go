package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"
	"encoding/json"
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

type device struct {
	DeviceID     uint16 `json:"device_id"`
	DeviceName   string `json:"device_name"`
	DeviceStatus bool   `json:"is_logging"`
	DeviceOwner  string `json:"device_owner"`
}

type CPUstats struct {
	CpuID   uint16  `json:"cpu_id"`
	CpuName string  `json:"cpu_name"`
	CpuFreq float64 `json:"cpu_freq"`
	CpuTemp float64 `json:"cpu_temp"`
}

type GPUstats struct {
	GpuID   uint16  `json:"gpu_id"`
	GpuName string  `json:"gpu_name"`
	GpuFreq float64 `json:"gpu_clock_speed"`
	GpuTemp float64 `json:"gpu_temp"`
}

type RAMstats struct {
	RamID    uint16  `json:"ram_id"`
	RamName  string  `json:"ram_name"`
	RamUsed  float64 `json:"ram_used"`
	RamTotal float64 `json:"ram_total"`
}

type DISKstats struct {
	DiskID   uint16  `json:"disk_id"`
	DiskName string  `json:"disk_name"`
	Disktemp float64 `json:"disk_temp"`
	DiskSize int     `json:"disk_size"`
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
		fmt.Println("Error loading .env file")
	}

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	ip := os.Getenv("DATABASE_IP")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	return LOGIN{
		User: user,
		Pass: pass,
		Ip:   ip,
		Port: port,
		Name: name,
	}
}

func setDeviceInformation() device {
	// Get the system information from the database

	newDevice := device{
		DeviceID:     1,
		DeviceName:   "",
		DeviceStatus: false,
	}
	return newDevice
}

func connectDB() sql.DB {

	credentials := loadCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.User, credentials.Pass, credentials.Ip, credentials.Port, credentials.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	showTables := "SHOW TABLES"
	rows, err := db.Query(showTables)
	if err != nil {
		fmt.Println("Error executing query:", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	fmt.Println("Tables in database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Println("Error scanning row:", err)
		}
		fmt.Println("- " + tableName)
	}
	fmt.Println("Connected to MySQL")
	return *db
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Server is running")
}

func UserInformation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "Welcome user, %s!\n", ps.ByName("name"))
}

func userReg(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user userInformation

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.UserPassword = string(hashedPassword)

	Insertion(user)

	w.WriteHeader(http.StatusCreated)
	err_ := json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user_id": user.UserID,
	})
	if err_ != nil {
		return
	}
}

func Insertion(data userInformation) {

	query := `INSERT INTO USER (user_name, user_password, user_role, user_email) 
              VALUES ($1, $2, $3, $4)`

	db := connectDB()
	err, _ := db.Exec(query, data.UserName, data.UserPassword, data.UserRole, data.UserEmail)
	if err != nil {
		log.Fatal("Error inserting data: ", err)
	} else {
		fmt.Println("User inserted successfully!")
	}
}

func serveServer() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", UserInformation)

	log.Fatal(http.ListenAndServe(":5210", router))
}

func main() {
	fmt.Println("getting system data...")

	connectDB()
	serveServer()
}
