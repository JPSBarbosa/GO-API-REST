package service

import (
	"github.com/JPSBarbosa/GO-API-REST/internal/models"
	"github.com/JPSBarbosa/GO-API-REST/internal/user"
)

type userService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &userService{repo: repo}
}

func (u *userService) FindById(id uint) (*models.User, error) {
	return u.repo.FindById(id)
}

func (u *userService) Save(user models.User) (uint, error) {
	return u.repo.Save(user)
}

func (u *userService) Update(user models.User) (*models.User, error) {
	return u.repo.Update(user)
}

func (u *userService) DeleteById(id uint) error {
	return u.repo.DeleteById(id)
}
