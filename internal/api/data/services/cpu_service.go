package services

import (
	"bufio"
	"log"
	"os/exec"
)

func Cpu(path string, parameter string) <-chan string {
	outChan := make(chan string)

	if parameter != "name" {
		cmd := exec.Command(path, parameter)
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
		return outChan, parameter
	}
}
