package logs

import (
	"fmt" 
	"os"
	"time"
)

func Log(){
	fmt.Println("logging...")
	currTime := []byte(time.Now().String())
	
	os.Create("logs/log.txt")
	os.WriteFile("log.txt", currTime, 0644 )
}

func LogErrors(err error) []byte {

	formattedError := []byte(fmt.Sprintf("%v", err))
	return formattedError 
}