package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var user, password, logLevel, podIp string

func init() {
	user = os.Getenv("APP_USER")
	password = os.Getenv("APP_PASSWORD")
	logLevel = os.Getenv("LOG_LEVEL")
	podIp = os.Getenv("POD_IP")
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("new request to %s", podIp))

	var response strings.Builder
	response.WriteString(fmt.Sprintf("USER: %s\n", user))
	response.WriteString(fmt.Sprintf("PASSWOR: %s\n", password))
	response.WriteString(fmt.Sprintf("LOG_LEVEL: %s\n", logLevel))
	response.WriteString(fmt.Sprintf("POD_IP: %s\n", podIp))
	response.WriteString(fmt.Sprintf("TIME: %v\n", time.Now()))

	_, err := w.Write([]byte(response.String()))
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	log.SetOutput(file)
	log.Println("app started")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
