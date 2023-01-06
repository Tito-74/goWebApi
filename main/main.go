package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func aboutMe(c *fiber.Ctx) error{
	fmt.Printf(
		"My name is Tito, am a software developer at dawascope technologies.Am working on backend where am tackling everthing concerning database, Api, and secuirity",
	)
	return nil
	
}

func homePage(c *fiber.Ctx) error {
	fmt.Printf("Welcome to the web API")
	fmt.Printf("Endpoint Hit:hompage")
	return nil
}
func main() {
	app := fiber.New()

	app.Get("/homepage", homePage)
	app.Get("/aboutMe", aboutMe)

	app.Listen(":3000")
}

