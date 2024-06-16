package server

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productController struct {
	service product.Service
}

func NewProductController(service product.Service) product.Controller {
	return &productController{service: service}
}

func (p *productController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p *productController) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := p.service.Save(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	c.JSON(http.StatusOK, product)
}

func (p *productController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)

	updatedProduct, err := p.service.Update(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := p.service.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
