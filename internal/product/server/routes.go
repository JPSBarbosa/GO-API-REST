package server

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/product"
	"github.com/gin-gonic/gin"
)

func MapProductRoutes(r *gin.Engine, controller product.Controller) {

	r.GET("/product/:id", controller.Get)
	r.POST("/product", controller.Create)
	r.PUT("/product/:id", controller.Update)
	r.DELETE("/product/:id", controller.Delete)
}
