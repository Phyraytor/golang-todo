package repository

import (
	"database/sql"
	models "github.com/Phyraytor/golang-todo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Authorization
	Foodstuff
	PostgresFoodstuff
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Foodstuff:     NewFoodstuffMysql(db),
	}
}
func NewRepositoryPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		PostgresFoodstuff: NewFoodstuffPostgres(db),
	}
}
