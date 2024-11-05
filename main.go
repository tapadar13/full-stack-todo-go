package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo
var idCounter = 1

func main() {
    app := fiber.New()

	// Create a new todo
	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = idCounter
		idCounter++
		todos = append(todos, *todo)
		return c.JSON(todo)
	})

	// Get all todos
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Get a specific todo
	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(todo)
			}
		}
		return c.Status(404).SendString("Todo not found")
	})

	// Update a todo
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		updatedTodo := new(Todo)
		if err := c.BodyParser(updatedTodo); err != nil {
			return err
		}
		for i, todo := range todos {
			if todo.ID == id {
				updatedTodo.ID = id
				todos[i] = *updatedTodo
				return c.JSON(updatedTodo)
			}
		}
		return c.Status(404).SendString("Todo not found")
	})

	// Delete a todo
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendString("Todo deleted")
			}
		}
		return c.Status(404).SendString("Todo not found")
	})

	log.Fatal(app.Listen(":4000"))
}