package main

import (
	"encoding/json"
	"fmt"
	"product/internal/config"
)

func main() {

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	str, err := json.Marshal(conf)

	fmt.Printf("Config: %s", string(str))
	//appName := product
	//appHost := localhost
	//appPort := 4447
	//mailFrom := test@test.ru
	//mailHost := smtp.gmail.com
	//mailPort := 587
	//mailSendTopic := send_message

}
