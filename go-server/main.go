package main

import (
	"fmt"
	"log"

	"net/http"
	"encoding/json"

	"database/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"os"
)

type userInformation struct
{
	userID uint16 `json:"id"`
	userName string `json:"user_name"` 
	userRole bool `json:"user_role"`
	userEmail string `json:"user_email"`
	userEmailVerified bool `json:"is_email_verified"`
}

type device struct {
	deviceID   uint16 `json:"device_id"` 
	deviceName string `json:"device_name"`
	deviceStatus bool `json:"is_loggin"`   
}

type CPUstats struct {
	cpuID	uint16 `json:cpu_id`	
	cpuName	string `json:"cpu_name"`
	cpuFreq	float64 `json:"cpu_clock_speed"` 
	cpuTemp float64
}

type GPUstats struct
{
	gpuID uint16 `json:"gpu_id"`
	gpuName string `json:"gpu_name"`
	gpuFreq float64 `json:"gpu_clock_speed"`
	gpuTemp float64 `json:"gpu_temp"`

}

type RAMstats struct
{
	ramID		uint64  `json:"ram_id"` 
	ramName      string `json:"ram_name"`
	ramUsed      float64 `json:"ram_used"`
	ramTotal     float64 `json:"ram_total"`
}

type DISKstats struct
{
	diskID uint64 `json:"disk_id"`
	diskName string `json:"disk_name"`
	disktemp float64 `json:"disk_temp"`
	diskSize int `json:"diskName"`
}

type LOGIN struct 
{
	user string
	pass string
	ip   string 
	port string
}

func loadCredentials() LOGIN{

	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	ip := os.Getenv("DATABASE_IP")
	port := os.Getenv("DATABASE_PORT")

	return LOGIN{
		user: user,
		pass: pass,
		ip: ip,
		port: port,
	}
}

func setDeviceInformation() device {
	// Get the system information from the database

	newDevice := device{
		deviceID: 1,
		deviceName: "", 
		deviceStatus: false,
	}
	return newDevice
}

func updateUserInformatio()  {


	sql.Open("mysql", "`{credentials.user}`:root@tcp(host:port)/dbname")
}


func main() {
	fmt.Println("getting system data...")
}
