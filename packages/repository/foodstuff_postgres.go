package repository

import (
	"fmt"
	models "github.com/Phyraytor/golang-todo/model"
	"github.com/jmoiron/sqlx"
)

type FoodstuffPostgres struct {
	db *sqlx.DB
}

func NewFoodstuffPostgres(db *sqlx.DB) *FoodstuffPostgres {
	return &FoodstuffPostgres{db: db}
}

func (r *FoodstuffPostgres) Create(foodstuff models.Foodstuff) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO foodstuff (Protein, Fat, Carb, Price) VALUES ($1, $2, $3, $4) RETURNING id")
	result, err := r.db.Exec(query, foodstuff.Protein, foodstuff.Fat, foodstuff.Carb, foodstuff.Price)
	if err != nil {
	} else {
		insertId, err := result.LastInsertId()
		if err == nil {
			id = int(insertId)
		}
	}
	return id, nil
}

func (r *FoodstuffPostgres) GetAll() ([]models.Foodstuff, error) {
	var items []models.Foodstuff
	query := fmt.Sprintf("SELECT * FROM foodstuff")
	err := r.db.Select(&items, query)
	return items, err
}

func (r *FoodstuffPostgres) GetById(id int) (models.Foodstuff, error) {
	var item models.Foodstuff
	err := r.db.Get(&item, "SELECT * FROM foodstuff WHERE id = ?", id)
	return item, err
}

func (r *FoodstuffPostgres) Update(id int, input models.UpdateFoodstuffInput) error {
	_, err := r.db.Exec("UPDATE foodstuff SET Protein = ?, Fat = ?, Carb = ?, Price = ? WHERE id = ?",
		input.Protein, input.Fat, input.Carb, input.Price, id)
	return err
}

func (r *FoodstuffPostgres) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM foodstuff WHERE id = ?", id)
	return err
}
