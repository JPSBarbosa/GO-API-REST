package product

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
}
