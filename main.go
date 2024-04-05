package main

import (
	"database/sql"
	"github.com/JPSBarbosa/GO-API-REST/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=postgres password=postgres dbname=GO-API-REST"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	routes.SetupSessionRoutes(app)

	app.Listen(":3000")
	defer db.Close()
}
