package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/prajkin/em-test-task/database"
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
	if err != nil {
		log.Fatal(err)
	}
}
