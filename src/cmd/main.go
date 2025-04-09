package main

import (
	"fmt"

	sys "synf/handlers/data"
	"synf/server"
)

func main() {
	// clear the terminal using ASCII code
	//	fmt.Println("\033[H\033[2J")

	sys.CpuTemperature()

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())
}
