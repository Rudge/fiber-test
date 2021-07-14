package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) { c.JSON(&Token{"access", "token"}) })

	app.Listen(":" + port)
}
