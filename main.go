package main

import (
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/prajkin/em-test-task/database"
	_ "github.com/prajkin/em-test-task/docs"
	"github.com/prajkin/em-test-task/internal/handler"
	"github.com/prajkin/em-test-task/internal/repository"
	"github.com/prajkin/em-test-task/internal/service"
)

// @title Subscriptions API
// @version 1.0
// @description Simple subscription service
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("db connected succesfully")

	err = database.MigrateUp(db)
	if err == migrate.ErrNoChange {
		log.Println("db migrations: no changes")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Println("db migrated succesfully")
	}

	repo := repository.NewSubscriptionsRepository(db)
	svc := service.NewSubscriptionsService(repo)
	h := handler.NewHandler(svc)
	if err = http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
