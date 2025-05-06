package main

import (
	"fmt"
	dataConnection "synf/internal/api/ws"
	"synf/internal/server"
)

func runAPI() {

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())

}
func getCPU() {

	fmt.Println("Pulling CPU Temperature")
	dataConnection.InitCPU()


}
func main() {

	fmt.Println("\033[H\033[2J")
	getCPU()
	runAPI()
	fmt.Println("Welcome to the terminal environment of synf")

}
