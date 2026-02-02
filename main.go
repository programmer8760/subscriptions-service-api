package main

import (
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/prajkin/em-test-task/database"
	"github.com/prajkin/em-test-task/internal/handler"
	"github.com/prajkin/em-test-task/internal/repository"
	"github.com/prajkin/em-test-task/internal/service"
)

func main() {
	godotenv.Load()

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
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	repo := repository.NewSubscriptionsRepository(db)
	svc := service.NewSubscriptionsService(repo)
	h := handler.NewHandler(svc)
	if err = http.ListenAndServe(":3000", h); err != nil {
		log.Fatal(err)
	}
}
