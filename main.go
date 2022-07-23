package main

import (
	"code.project.com/InstantMessaging/models"
	"code.project.com/InstantMessaging/router"
	"fmt"
)

func main() {
	engine := router.Router()
	models.MongoInit()
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
