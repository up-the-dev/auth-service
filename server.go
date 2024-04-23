package main

import (
	"fmt"
	application "github.com/up-the-dev/auth-service/app"
	"github.com/up-the-dev/auth-service/config"
	"os"
)

func main() {
	fmt.Println("Welcome to auth-svc")

	// getting instance of fiber application. we created instance of app so that we can use that instance for testing also
	app := application.Get_instance()
	// loading environment variables
	configuration_error := config.Load()
	if configuration_error != nil {
		fmt.Println("error loading configurations : ", configuration_error)
		os.Exit(1)
	}

	//listening service
	error := app.Listen(":" + os.Getenv("APP_PORT"))
	if error != nil {
		fmt.Println("error starting server : ", error)
		os.Exit(1)
	}
}
