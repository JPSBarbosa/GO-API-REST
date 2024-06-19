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

// Get godoc
// @Summary Get a product by ID
// @Description Get product by ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string][]string
// @Router /product/{id} [get]
func (p *productController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Product to create"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /product [post]
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

// Update godoc
// @Summary Update an existing product
// @Description Update a product by ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product data to update"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /product/{id} [put]
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

// Delete godoc
// @Summary Delete a product by ID
// @Description Delete product by ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /product/{id} [delete]
func (p *productController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := p.service.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
