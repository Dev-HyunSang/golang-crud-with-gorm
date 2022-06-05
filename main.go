package main

import (
	"log"

	"github.com/dev-hyunsang/golang-todo-list/database"
	"github.com/dev-hyunsang/golang-todo-list/middleware"
	"github.com/dev-hyunsang/golang-todo-list/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Router(app)

	db, err := database.ConnectionDB()
	if err != nil {
		log.Println("Failed to Connection DataBase")
		log.Println(err)
	}

	db.Table("todos").AutoMigrate(&models.ToDo{})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to Listen Server %v\n", err)
	}
}
