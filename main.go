package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1540877",
		Key:     "1d82560ac6cbb9fdec3f",
		Secret:  "2daee592d2a3c9b2817c",
		Cluster: "ap2",
		Secure:  true,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		err := pusherClient.Trigger("chat", "message", data)
		if err != nil {
			fmt.Println(err.Error())
		}
		return c.JSON([]string{})
	})

	app.Listen(":3000")
}
