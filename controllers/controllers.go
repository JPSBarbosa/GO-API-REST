package controllers

import (
	"github.com/JPSBarbosa/GO-API-REST/Models"
	"github.com/JPSBarbosa/GO-API-REST/Service"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func GetSession(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(id)
}

func Homepage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(`
	<html>
		<head>
			<title>Página Inicial</title>
		</head>
		<body>
			<button onclick="window.location.href='/Teste'">Ir para página teste</button>
		</body>
	</html>`)
}

func Teste(c *fiber.Ctx) error {
	return c.SendString("Teste")
}

type ProductController struct {
	Service *Service.ProductsService
}

func NewProductController(service *Service.ProductsService) *ProductController {
	return &ProductController{Service: service}
}

func (p *ProductController) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.Service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product Models.Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := p.Service.SaveProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	c.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Models.Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := p.Service.UpdateProduct(product, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := p.Service.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
