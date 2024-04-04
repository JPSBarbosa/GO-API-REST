package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New()

	app.Get("/session/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(id)
	})

	app.Listen(":3000")
}
