package service

import (
	models "github.com/Phyraytor/golang-todo/model"
	"github.com/Phyraytor/golang-todo/packages/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Foodstuff interface {
	CreateFoodstuff(foodstuff models.Foodstuff) (int, error)
	GetAllFoodstuff() ([]models.Foodstuff, error)
	GetFoodstuffById(id int) (models.Foodstuff, error)
	UpdateFoodstuff(id int, input models.UpdateFoodstuffInput) error
	DeleteFoodstuff(id int) error
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Foodstuff
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Foodstuff:     NewFoodstuffService(repos.Foodstuff),
	}
}
