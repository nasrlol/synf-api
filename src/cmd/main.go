package main

import (
	"fmt"

	//	sys "synf/handlers/data"
	dataConnection "synf/handlers/connect"
	"synf/server"
)

func main() {
	fmt.Println("\033[H\033[2J")

	fmt.Println("Pulling CPU Temperature")
	dataConnection.Init()

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())
}
