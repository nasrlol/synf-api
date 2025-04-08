package sys

import (
	"fmt"
	"os/exec"
)

func CpuTemperature() string {
	fmt.Println("Getting CPU temperature")
	cpuTemperature, err := exec.Command("../../../synf-sys/cpu", "temperature").Output()
	if err != nil{
		fmt.Println("can't get cpu temperature")
	}
	return string(cpuTemperature);
}

func CpuName() string {
	fmt.Println("Getitng the CPU Name")
	cpuName, err := exec.Command("/synf-sys/cpu", "name").Output()
	if err != nil{
		fmt.Println("can't get the cpu name")
	}
	return string(cpuName)
}

func CpuFrequency() string {
	fmt.Println("Getting the CPU frequency")
	cpuFreq, err := exec.Command("/synf-sys/cpu", "frequency").Output()
		if err != nil{
		fmt.Println("can't get the cpu name")
	}
	return string(cpuFreq)
}

func Sys(){
	CpuName()
}
