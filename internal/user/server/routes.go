package server

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/user"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(t *gin.Engine, Controller user.Controller) {
	t.GET("/user/:id", Controller.Get)
	t.PUT("/user/:id", Controller.Update)
	t.DELETE("/user/:id", Controller.Delete)
	t.POST("/user", Controller.Create)
}
