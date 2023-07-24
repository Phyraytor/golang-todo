package repository

import (
	"database/sql"
	models "github.com/Phyraytor/golang-todo/model"
	_ "github.com/go-sql-driver/mysql"
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
type Repository struct {
	Authorization
	Foodstuff
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Foodstuff:     NewFoodstuffMysql(db),
	}
}
