package main

import (
	"fmt"
	"user-service/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	application.Router.Run(":8080")
}
