package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	connectionString := "user=postgres password=postgres dbname=GO-API-REST"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/session/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(id)
	})

	app.Listen(":3000")
	defer db.Close()
}
