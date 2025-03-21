package main

import (
    "fmt"
)

type device struct {
	device_id   int
	device_name string
}

type stats struct {
	cpu_frequency float64
	cpu_temp      float64
	cpu_name      string

    ram_name      string
    ram_used      int
    ram_free      int
    ram_total     int

}

func getSystemInformation() device {
	return device{}
}

func main() {
	newCpu := getSystemInformation()
	fmt.Println(newCpu)
}
