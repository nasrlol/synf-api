package cpu

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

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
