package service

import (
	models "github.com/Phyraytor/golang-todo/model"
	"github.com/Phyraytor/golang-todo/packages/repository"
)

type FoodstuffService struct {
	repo repository.Foodstuff
}

func NewFoodstuffService(repo repository.Foodstuff) *FoodstuffService {
	return &FoodstuffService{repo: repo}
}

func (s *FoodstuffService) CreateFoodstuff(foodstuff models.Foodstuff) (int, error) {
	return s.repo.CreateFoodstuff(foodstuff)
}

func (s *FoodstuffService) GetAllFoodstuff() ([]models.Foodstuff, error) {
	return s.repo.GetAllFoodstuff()
}

func (s *FoodstuffService) GetFoodstuffById(id int) (models.Foodstuff, error) {
	return s.repo.GetFoodstuffById(id)
}

func (s *FoodstuffService) UpdateFoodstuff(id int, input models.UpdateFoodstuffInput) error {
	return s.repo.UpdateFoodstuff(id, input)
}

func (s *FoodstuffService) DeleteFoodstuff(id int) error {
	return s.repo.DeleteFoodstuff(id)
}
