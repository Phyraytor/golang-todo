package repository

import (
	"database/sql"
	"fmt"
	models "github.com/Phyraytor/golang-todo/model"
	"strconv"
)

type TFoodstuff struct {
	db *sql.DB
}

func NewFoodstuffMysql(db *sql.DB) *TFoodstuff {
	return &TFoodstuff{db: db}
}

func (r *TFoodstuff) CreateFoodstuff(foodstuff models.Foodstuff) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO foodstuff (Protein, Fat, Carb, Price) values (?, ?, ?, ?)")
	result, err := r.db.Exec(query, foodstuff.Protein, foodstuff.Fat, foodstuff.Carb, foodstuff.Price)
	if err != nil {
	} else {
		insertId, err := result.LastInsertId()
		if err == nil {
			fmt.Printf(strconv.FormatInt(insertId, 10))
			id = int(insertId)
		}
	}
	return id, nil
}

func (r *TFoodstuff) GetAllFoodstuff() ([]models.Foodstuff, error) {
	rows, err := r.db.Query("SELECT * FROM foodstuff WHERE 1")
	if err != nil {
	}
	foodstuffs := []models.Foodstuff{}
	for rows.Next() {
		p := models.Foodstuff{}
		err := rows.Scan(&p.Id, &p.Protein, &p.Fat, &p.Carb, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		foodstuffs = append(foodstuffs, p)
	}
	return foodstuffs, nil
}

func (r *TFoodstuff) GetFoodstuffById(id int) (models.Foodstuff, error) {
	result := r.db.QueryRow("SELECT * FROM foodstuff WHERE id = ?", id)
	p := models.Foodstuff{}
	err := result.Scan(&p.Id, &p.Protein, &p.Fat, &p.Carb, &p.Price)
	if err != nil {
	}
	return p, nil
}

func (r *TFoodstuff) UpdateFoodstuff(id int, input models.UpdateFoodstuffInput) error {
	_, err := r.db.Exec("UPDATE foodstuff SET Protein = ?, Fat = ?, Carb = ?, Price = ? WHERE id = ?",
		input.Protein, input.Fat, input.Carb, input.Price, id)
	return err
}

func (r *TFoodstuff) DeleteFoodstuff(id int) error {
	_, err := r.db.Exec("DELETE FROM foodstuff WHERE id = ?", id)
	return err
}
