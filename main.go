package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/prajkin/em-test-task/database"
	_ "github.com/prajkin/em-test-task/docs"
	"github.com/prajkin/em-test-task/internal/handler"
	"github.com/prajkin/em-test-task/internal/logger"
	"github.com/prajkin/em-test-task/internal/repository"
	"github.com/prajkin/em-test-task/internal/service"
)

// @title Subscriptions API
// @version 1.0
// @description Simple subscription service
// @host localhost:8080
// @BasePath /
func main() {
	log := logger.New(slog.LevelInfo)

	if _, ok := os.LookupEnv("HTTP_PORT"); !ok {
		err := godotenv.Load()
		if err != nil {
			log.Error("loading .env file failed", "err", err)
			os.Exit(1)
		}
	}

	db, err := database.Connect()
	if err != nil {
		log.Error("database connection failed", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Error("database didn't respond to ping", "err", err)
		os.Exit(1)
	}
	log.Info("database connection established")

	err = database.MigrateUp(db)
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info("database migrations: no changes")
		} else {
			log.Error("database migration failed", "err", err)
			os.Exit(1)
		}
	} else {
		log.Info("database migrations: success")
	}

	repo := repository.NewSubscriptionsRepository(db)
	svc := service.NewSubscriptionsService(repo, log)
	h := handler.NewHandler(svc, log)

	addr := ":" + os.Getenv("HTTP_PORT")
	log.Info("http server started", "addr", addr)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Error("http server stopped", "err", err, "addr", addr)
		os.Exit(1)
	}
}
