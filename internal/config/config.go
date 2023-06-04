package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	defaultAppName       = "product"
	defaultAppHost       = "localhost"
	defaultAppPort       = "4447"
	defaultMailFrom      = "test@test.ru"
	defaultMailHost      = "smtp.gmail.com"
	defaultMailPort      = "587"
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
	MailSend *MailSendConfig `json:"mail_send"`
}

type Config struct {
	App   *AppConfig   `json:"app"`
	Mail  *MailConfig  `json:"mail"`
	Event *EventConfig `json:"event"`
}

type MailSendConfig struct {
	Topic string `json:"topic"`
}

func New() (*Config, error) {

	var err error

	appName, ok := os.LookupEnv("APP_NAME")
	if !ok {
		appName = defaultAppName
		err = os.Setenv("APP_NAME", appName)
	}

	appHost, ok := os.LookupEnv("APP_HOST")
	if !ok {
		appHost = defaultAppHost
		err = os.Setenv("APP_HOST", appHost)
	}

	appPort, ok := os.LookupEnv("APP_PORT")
	if !ok {
		appPort = defaultAppPort
		err = os.Setenv("APP_PORT", appPort)
	}

	mailFrom, ok := os.LookupEnv("MAIL_FROM")
	if !ok {
		mailFrom = defaultMailFrom
		err = os.Setenv("MAIL_FROM", mailFrom)
	}

	mailHost, ok := os.LookupEnv("MAIL_HOST")
	if !ok {
		mailHost = defaultMailHost
		err = os.Setenv("MAIL_HOST", mailHost)
	}

	mailPort, ok := os.LookupEnv("MAIL_PORT")
	if !ok {
		mailPort = defaultMailPort
		err = os.Setenv("MAIL_PORT", mailPort)
	}

	mailSendTopic, ok := os.LookupEnv("MAIL_SEND_TOPIC")
	if !ok {
		mailSendTopic = defaultMailSendTopic
		err = os.Setenv("MAIL_SEND_TOPIC", mailSendTopic)
	}

	jsonBytes, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Printf("Open config error: %s", err.Error())
		return nil, err
	}
	jsonString := os.ExpandEnv(string(jsonBytes))
	var config *Config
	err = json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		fmt.Printf("Unmarshaling error: %s", err.Error())
		return nil, err
	}

	return config, nil
}
