package sys


import (
	"fmt"
	"os/exec"
)

func Sys(){
	fmt.Println("getting system data")
	cpuTemperature, err := exec.Command("../../../synf-sys/cpu").Output()
	if err != nil{
		fmt.Println("cant get cpu information")
		return
	}
	fmt.Println(cpuTemperature)
}
