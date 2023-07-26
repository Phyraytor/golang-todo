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

type PostgresFoodstuff interface {
	Create(foodstuff models.Foodstuff) (int, error)
	GetAll() ([]models.Foodstuff, error)
	GetById(id int) (models.Foodstuff, error)
	Update(id int, input models.UpdateFoodstuffInput) error
	Delete(id int) error
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Foodstuff
	PostgresFoodstuff
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:     NewAuthService(repos.Authorization),
		PostgresFoodstuff: NewFoodstuffService(repos.PostgresFoodstuff),
	}
}
