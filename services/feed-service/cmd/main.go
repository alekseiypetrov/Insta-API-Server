package main

import (
	"fmt"
	"project/services/feed-service/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	application.Router.Run(":8080")
}
