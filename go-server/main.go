package main

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

type userInformation struct {
	UserID            uint16 `json:"id"`
	UserName          string `json:"user_name"`
	UserRole          bool   `json:"user_role"`
	UserEmail         string `json:"user_email"`
	UserEmailVerified bool   `json:"is_email_verified"`
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
	RamID    uint64  `json:"ram_id"`
	RamName  string  `json:"ram_name"`
	RamUsed  float64 `json:"ram_used"`
	RamTotal float64 `json:"ram_total"`
}

type DISKstats struct {
	DiskID   uint64  `json:"disk_id"`
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

func connectDB() {

	credentials := loadCredentials()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.User, credentials.Pass, credentials.Ip, credentials.Port, credentials.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
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
		return
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
			return
		}
		fmt.Println("- " + tableName)
	}
	fmt.Println("Connected to MySQL")
}

func main() {
	fmt.Println("getting system data...")
	connectDB()
}
