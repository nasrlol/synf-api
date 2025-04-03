package logs

import (
	"fmt" 
	"os"
	"time"
)

func logErrors(err error) []byte {
	formattedError := []byte(fmt.Sprintf("%v", err))
	return formattedError 
}

func Log(err error){
	fmt.Println("logging...")
	currTime := []byte(time.Now().String())
	
	os.Create("logs/log.txt")
	os.WriteFile("log.txt", currTime, 0644 )
	os.WriteFile("log.txt", logErrors(err), 0644)
}