package main

import (
	_ "github.com/JPSBarbosa/GO-API-REST/docs"
	productRepository "github.com/JPSBarbosa/GO-API-REST/internal/product/repository"
	productServer "github.com/JPSBarbosa/GO-API-REST/internal/product/server"
	productService "github.com/JPSBarbosa/GO-API-REST/internal/product/service"
	userRepository "github.com/JPSBarbosa/GO-API-REST/internal/user/repository"
	userServer "github.com/JPSBarbosa/GO-API-REST/internal/user/server"
	userService "github.com/JPSBarbosa/GO-API-REST/internal/user/service"
	"github.com/JPSBarbosa/GO-API-REST/pkg/database"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO-API-REST
// @version 1.0
// @description API REST in GO.
// @host localhost:8080
// @scheme https
// @BasePath /
// swagger:meta
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
