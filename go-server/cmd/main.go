package main

import (
	"fmt"
	"synf/server"
)

func main() {

	// clear terminal screen ASCII sequence
	fmt.Print("\033[H\033[2J")

	server.RawConnect("127.0.0.1", "5000")
	server.RegistrationEndpoint()
	fmt.Println("API STARTED...")
	fmt.Println(server.GetOutboundIp())
	
}
