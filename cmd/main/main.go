package main

import (
	"product/internal/app"
	"product/internal/config"
)

func main() {

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	appInstance := app.New(conf.App)
	appInstance.Run()

}
