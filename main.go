package main

import (
	"fmt"

	"github.com/baristikir/go-fiber/book"
	"github.com/baristikir/go-fiber/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	
	if err != nil {
		panic("Failed to connect to database ")
	}
	
	fmt.Println("Database connection successfully opened")
	
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	initDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}