package main

import (
	"code.project.com/InstantMessaging/models"
	"code.project.com/InstantMessaging/pkg/config"
	"code.project.com/InstantMessaging/router"
	"fmt"
)

func main() {
	err := config.LoadConfig("./config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	engine := router.Router()
	models.MongoInit()
	err = engine.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
