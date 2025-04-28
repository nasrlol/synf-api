package main

import (
	"fmt"
	dataConnection "synf/api/ws"

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
		fmt.Errorf("couldn't run npm")
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
	go getCPU()
	go runAPI()
	fmt.Println("Welcome to the terminal environment of synf")

}
