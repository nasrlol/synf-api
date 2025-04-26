package main

import (
	"fmt"
	dataConnection "synf/handlers/connection/dataWebsocket"

	"os"
	"os/exec"
	"synf/server"
)

func runNpmDev() {

	cmd := exec.Command("npm", "run", "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Errorf("coudlnt run npm")
	}
}

func runAPI() {

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())

}
func getCPU() {

	fmt.Println("Pulling CPU Temperature")
	dataConnection.Init()

}
func main() {

	fmt.Println("\033[H\033[2J")
	getCPU()
	go runAPI()
	fmt.Println("Welcome to the terminal envirment of SYNF")

}
