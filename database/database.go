package database

import (
	"log"
	"time"

	"github.com/dev-hyunsang/golang-todo-list/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	return db, err
}

func CreateToDo(todos *models.ToDo) *gorm.DB {
	db, err := ConnectionDB()
	if err != nil {
		log.Println("Failed to Connection DataBase via SQLite")
		log.Println(err)
	}

	result := db.Table("todos").Create(todos)

	return result
}

func ReadToDo() []models.ToDo {
	db, err := ConnectionDB()
	if err != nil {
		log.Println("Failed to Connection DataBase via SQLite")
		log.Println(err)
	}

	var resultToDo []models.ToDo
	db.Table("todos").Find(&resultToDo)

	return resultToDo
}

func EditToDo(todoUUID, content string) {
	db, err := ConnectionDB()
	if err != err {
		log.Println("Failed to Connection DataBase via SQLite")
		log.Println(err)
	}

	db.Table("todos").Where("todo_uuid = ?", todoUUID).Update("content", content)
	db.Table("todos").Where("todo_uuid = ?", todoUUID).Update("updated_time", time.Now())
}

func DeleteToDo(todoUUID string) {
	db, err := ConnectionDB()
	if err != nil {
		log.Println("Failed to Connection DataBase via SQLite")
		log.Println(err)
	}

	db.Table("todos").Where("todo_uuid = ?", todoUUID).Delete(&models.ToDo{})
}
