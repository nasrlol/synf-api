package sys

import (
	"fmt"
	"os/exec"
)

func CpuTemperature(){
	fmt.Println("Getting CPU temperature")
	cpuTemperature, err := exec.Command("../../../synf-sys/cpu", "temp").Output()
	if err != nil{
		fmt.Println("can't get cpu temperature")
		return
	}
	fmt.Println(cpuTemperature)
}

func CpuName(){
	fmt.Println("Getitng the CPU Name")
	cpuName, err := exec.Command("/synf-sys/cpu", "name").Output()
	if err != nil{
		fmt.Println("can't get the cpu name")
	}
	fmt.Println(cpuName)
}

func CpuFrequency(){
	fmt.Println("Getting the CPU frequency")
	cpuFreq, err := exec.Command("/synf-sys/cpu", "freq").Output()
		if err != nil{
		fmt.Println("can't get the cpu name")
	}
	fmt.Println(cpuFreq)
}

func Sys(){
	CpuTemperature()
	CpuName()
}