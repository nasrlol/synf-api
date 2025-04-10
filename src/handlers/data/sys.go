package sys

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func CpuTemperature() {
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

func dataPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "data")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "websocket endpoint")
}

func setupRouters() {
	http.HandleFunc("/", dataPage)
	http.HandleFunc("ws", wsEndpoint)
}

func Run() {
	fmt.Println("Websocket")
	setupRouters()
	log.Fatal(http.ListenAndServe(":8090", nil))
}
