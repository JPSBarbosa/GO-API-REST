package main

import (
	productRepository "github.com/JPSBarbosa/GO-API-REST/internal/product/repository"
	productServer "github.com/JPSBarbosa/GO-API-REST/internal/product/server"
	productService "github.com/JPSBarbosa/GO-API-REST/internal/product/service"
	userRepository "github.com/JPSBarbosa/GO-API-REST/internal/user/repository"
	userServer "github.com/JPSBarbosa/GO-API-REST/internal/user/server"
	userService "github.com/JPSBarbosa/GO-API-REST/internal/user/service"
	"github.com/JPSBarbosa/GO-API-REST/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	psqlConn := database.PostgreSQL()
	productRepo := productRepository.NewProductRepository(psqlConn)
	productService := productService.NewProductService(productRepo)
	productController := productServer.NewProductController(productService)

	productServer.MapProductRoutes(r, productController)

	userRepo := userRepository.NewUserRepository(psqlConn)
	uService := userService.NewUserService(userRepo)
	userController := userServer.NewUserController(uService)

	userServer.MapUserRoutes(r, userController)

	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
