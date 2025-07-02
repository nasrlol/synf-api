package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func UserLog(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("endpoint is running, user login endpoint"))
}

func RegistrationEndpoint() {
	router := httprouter.New()
	router.GET("/user/login", registration.GetUser)
	router.POST("/user/registration", registration.UserReg)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func RawConnect(host string, port string) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connecting error:", err)
	}
	if conn != nil {
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
		fmt.Println("Opened", net.JoinHostPort(host, port))
	}
}
