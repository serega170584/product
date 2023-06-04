package main

import (
	"encoding/json"
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

type AppConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type MailConfig struct {
	From string `json:"from"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type EventConfig struct {
	MailSend MailSendConfig `json:"mail_send"`
}

type MailSendConfig struct {
	Topic string `json:"topic"`
}

type Config struct {
	App   AppConfig   `json:"app"`
	Mail  MailConfig  `json:"mail"`
	Event EventConfig `json:"event"`
}

func main() {
	jsonBytes, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Printf("Open config error: %s", err.Error())
	}
	jsonString := os.ExpandEnv(string(jsonBytes))
	var config *Config
	err = json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		fmt.Printf("Unmarshaling error: %s", err.Error())
	}
	fmt.Printf("Topic: %s", config.Event.MailSend.Topic)
	//appName := product
	//appHost := localhost
	//appPort := 4447
	//mailFrom := test@test.ru
	//mailHost := smtp.gmail.com
	//mailPort := 587
	//mailSendTopic := send_message

}
