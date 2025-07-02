package services

import (
	"bufio"
	"log"
	"os/exec"
)

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
