package main

import (
	"fmt"
	"os"
)

const (
	defaultAppName       = "product"
	defaultAppHost       = "localhost"
	defaultAppPort       = 4447
	defaultMailFrom      = "test@test.ru"
	defaultMailHost      = "smtp.gmail.com"
	defaultMailPort      = 587
	defaultMailSendTopic = "send_message"
)

func main() {
	jsonString, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Printf("Open config error: %s", err.Error())
	}
	fmt.Print(string(jsonString))
	//appName := product
	//appHost := localhost
	//appPort := 4447
	//mailFrom := test@test.ru
	//mailHost := smtp.gmail.com
	//mailPort := 587
	//mailSendTopic := send_message

}
