package main

import (
	"github.com/JPSBarbosa/GO-API-REST/Service"
	"github.com/JPSBarbosa/GO-API-REST/controllers"
	"github.com/JPSBarbosa/GO-API-REST/database"
	"github.com/JPSBarbosa/GO-API-REST/routes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	database.Init()
	db := database.DB

	app := fiber.New()

	routes.SetupSessionRoutes(app)

	productService := Service.NewProductService(db)
	productController := controllers.NewProductController(productService)

	r.GET("/products/:id", productController.GetProduct)
	r.POST("/products", productController.CreateProduct)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	app.Listen(":3000")
}
