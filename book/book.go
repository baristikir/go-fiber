package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All books")
}
func GetBook(c *fiber.Ctx) error {
	return c.SendString("A single book")
}
func NewBook(c *fiber.Ctx) error {
	return c.SendString("Adds a new book")
}
func DeleteBook(c *fiber.Ctx) error{
	return c.SendString("Deletes a book")
}