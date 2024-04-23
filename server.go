package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/up-the-dev/auth-service/config"
)

func main() {
	fmt.Println("Welcome to auth-svc")
	app := fiber.New()
	// loading environment variables
	configuration_error := config.Load()
	if configuration_error != nil {
		fmt.Println("error loading configurations : ", configuration_error)
		os.Exit(1)
	}

	error := app.Listen(":" + os.Getenv("APP_PORT"))
	if error != nil {
		fmt.Println("error starting server : ", error)
	}
}
