package middleware

import (
	"github.com/dev-hyunsang/golang-todo-list/cmd"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/create", cmd.CreateToDo)
	app.Post("/read", cmd.ReadToDo)
	app.Post("/edit", cmd.EditToDo)
	app.Delete("/delete/:uuid", cmd.DeleteToDo)
}
