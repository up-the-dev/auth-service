package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Welcome to auth-svc")

	app := fiber.New()

	app.Listen(":3000")
	if error != nil {
		fmt.Println("error starting server : ", error)
	} else {
		fmt.Println("listing on port 3000")
	}
}
