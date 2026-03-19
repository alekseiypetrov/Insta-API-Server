package main

import (
	"fmt"
	"user-service/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	app.Router.Run(":8080")
}
