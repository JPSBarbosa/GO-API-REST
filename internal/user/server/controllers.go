package server

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userController struct {
	service user.Service
}

func NewUserController(service user.Service) user.Controller {
	return &userController{service: service}
}

// Create godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Register"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /user [post]
func (u *userController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := u.service.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete a user by ID
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /user/{id} [delete]
func (u *userController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u.service.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted!"})
}

// Update godoc
// @Summary Update an existing user
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param product body models.User true "Product data to update"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /user/{id} [put]
func (u *userController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data models.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.ID = uint(id)

	updatedUser, err := u.service.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// Get godoc
// @Summary Get a user by ID
// @Description Get user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string][]string
// @Router /user/{id} [get]
func (u *userController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := u.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}
