package sys

import (
	"fmt"
	"os/exec"
)

func CpuTemperature() string {
	fmt.Println("getting CPU temperature")
	path := "./cpu"
	argument := "temperature"

	cpuTemperature, err := exec.Command(path, argument).Output()
	if err != nil {
		fmt.Println("can't get cpu temperature")
	}
	fmt.Println("passed program execution?")
	return string(cpuTemperature)
}

func CpuName() string {
	fmt.Println("getting the CPU Name")
	path := "./cpu"
	argument := "name"

	cpuName, err := exec.Command(path, argument).Output()
	if err != nil {
		fmt.Println("can't get the cpu name")
	}
	fmt.Println("passed program execution?")
	return string(cpuName)
}

func CpuFrequency() string {
	fmt.Println("getting the CPU frequency")
	path := "./cpu"
	argument := "frequency"

	cpuFreq, err := exec.Command(path, argument).Output()
	if err != nil {
		fmt.Println("can't get the cpu name")
	}
	return string(cpuFreq)
}

