package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/golang-todo-list/database"
	"github.com/dev-hyunsang/golang-todo-list/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func CreateToDo(c *fiber.Ctx) error {
	todos := new(models.ToDo)
	if err := c.BodyParser(todos); err != nil {
		log.Fatalf("Failed to BodyParser %v\n", err)
	}

	todoUUID, err := uuid.NewV4()
	if err != nil {
		log.Println("Failed to Create UUID v4")
		log.Println(err)
	}

	result := database.CreateToDo(&models.ToDo{
		ToDoUUID:    todoUUID,
		Content:     todos.Content,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})

	log.Println(result)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":        500,
			"message":       "새로운 ToDo를 생성하던 도중 오류가 발생했어요.",
			"error_message": err,
			"time":          time.Now(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "새로운 할일 항목을 만들었어요!",
		"time":    time.Now(),
	})
}

func ReadToDo(c *fiber.Ctx) error {
	result := database.ReadToDo()

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "성공적으로 생성되어 있는 항목들을 가지고 왔어요!",
		"datas":   result,
		"time":    time.Now(),
	})
}

func EditToDo(c *fiber.Ctx) error {
	todos := new(models.ToDo)
	if err := c.BodyParser(todos); err != nil {
		log.Fatalf("Failed to BodyParser %v\n", err)
	}

	database.EditToDo(todos.ToDoUUID.String(), todos.Content)
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "성공적으로 수정했어요!",
		"time":    time.Now(),
	})
}

func DeleteToDo(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	database.DeleteToDo(uuid)

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "성공적으로 할일 항목을 삭제했어요.",
		"time":    time.Now(),
	})
}
