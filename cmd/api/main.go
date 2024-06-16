package main

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/product/repository"
	"github.com/JPSBarbosa/GO-API-REST/internal/product/server"
	"github.com/JPSBarbosa/GO-API-REST/internal/product/service"
	"github.com/JPSBarbosa/GO-API-REST/pkg/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	psqlConn := database.PostgreSQL()
	productRepo := repository.NewProductRepository(psqlConn)
	productService := service.NewProductService(productRepo)
	productController := server.NewProductController(productService)

	server.MapProductRoutes(r, productController)

	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
