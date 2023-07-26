package main

import (
	"database/sql"
	"github.com/Phyraytor/golang-todo/packages/handler"
	"github.com/Phyraytor/golang-todo/packages/repository"
	"github.com/Phyraytor/golang-todo/packages/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

var database *sql.DB

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//runSever := new(todo.Server)

	/*	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/todo")
		database = db*/

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "root",
		DBName:   "todo",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Println(err)
	}
	repos := repository.NewRepositoryPostgres(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	handlers.InitRoutes().Run(":8001")
	/*	if err := runSever.Run("8001", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}*/
}
