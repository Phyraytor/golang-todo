package repository

import (
	"database/sql"
	"fmt"
	models "github.com/Phyraytor/golang-todo/model"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (name, username, password) values (?, ?, ?)")
	result, err := r.db.Exec(query, user.Name, user.Username, user.Password)
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

/*func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
*/
