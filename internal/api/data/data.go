package data

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

type DeviceInformation struct {
	DeviceID uint8  `json:"device_id"`
	UpTime   string `json:"device_upTime"`
}

type CpuInformation struct {
	CpuID   uint8  `json:"cpu_id"`
	CpuName string `json:"cpu_name"`
	CpuTemp uint8  `json:"cpu_temp"`
	CpuFreq uint8  `json:"cpu_freq"`
}

type GpuInformation struct {
	GpuID   uint8  `json:"cpu_id"`
	GpuName string `json:"cpu_name"`
	GpuTemp uint8  `json:"cpu_temp"`
	GpuFreq uint8  `json:"cpu_freq"`
}

type RamInformation struct {
	RamID   uint8  `json:"cpu_id"`
	RamName string `json:"cpu_name"`
	RamTemp uint8  `json:"cpu_temp"`
	RamFreq uint8  `json:"cpu_freq"`
}

type DiskInformation struct {
	DiskID    uint8  `json:"disk_id"`
	DiskName  string `json:"disk_name"`
	DiskTemp  uint8  `json:"disk_temp"`
	DiskSpeed uint8  `json:"disk_speed"`
}

func DeviceUpTime() <-chan string {
	outchan := make(chan string)

	cmd := exec.Command("./general")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		defer close(outchan)
		for scanner.Scan() {

			outchan <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Printf("scanner error %v\n", err)
		}
	}()

	return outchan
}

func CpuTemperature() <-chan string {
	outChan := make(chan string)

	cmd := exec.Command("./cpu", "temperature")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	go func() {
		defer close(outChan)
		for scanner.Scan() {
			outChan <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Printf("scanner error %v\n", err)
		}
	}()
	return outChan
}

func CpuTemperaturePrint() {
	cmd := exec.Command("./cpu", "temperature")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("scanner error:", err)
		}
	}()
	select {}
}

func CpuName() string {
	fmt.Println("getting the CPU name")
	path := "./cpu"
	argument := "name"

	cpuName, err := exec.Command(path, argument).Output()
	if err != nil {
		fmt.Println("can't get the cpu name")
	}
	return string(cpuName)
}

func CpuFrequency() string {
	fmt.Println("getting the CPU frequency")
	path := "./cpu"
	argument := "frequency"

	cpuFreq, err := exec.Command(path, argument).Output()
	if err != nil {
		fmt.Println("can't get the cpu frequency")
	}
	return string(cpuFreq)
}
