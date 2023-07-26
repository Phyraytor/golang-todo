package service

import (
	models "github.com/Phyraytor/golang-todo/model"
	"github.com/Phyraytor/golang-todo/packages/repository"
)

type FoodstuffService struct {
	repo repository.PostgresFoodstuff
}

func NewFoodstuffService(repo repository.PostgresFoodstuff) *FoodstuffService {
	return &FoodstuffService{repo: repo}
}

func (s *FoodstuffService) Create(foodstuff models.Foodstuff) (int, error) {
	return s.repo.Create(foodstuff)
}

func (s *FoodstuffService) GetAll() ([]models.Foodstuff, error) {
	return s.repo.GetAll()
}

func (s *FoodstuffService) GetById(id int) (models.Foodstuff, error) {
	return s.repo.GetById(id)
}

func (s *FoodstuffService) Update(id int, input models.UpdateFoodstuffInput) error {
	return s.repo.Update(id, input)
}

func (s *FoodstuffService) Delete(id int) error {
	return s.repo.Delete(id)
}
