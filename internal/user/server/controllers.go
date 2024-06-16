package server

import "C"
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

func (u *userController) Create(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := u.service.FindById(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (u *userController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u.service.DeleteById(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted!"})
}

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

func (u *userController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := u.service.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}
