package book

import (
	"bookManagement/database"
	"bookManagement/model"
	"bookManagement/validator"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book validator.BookData

	err := c.BodyParser(&book)
	if err != nil {
		c.Status(500)
		c.SendString(err.Error())
		return
	}
	errorValidation := book.Validate()
	if errorValidation != nil {
		c.Status(500).JSON(errorValidation)
	} else {
		var b model.Book
		b.Title = book.Title
		b.Author = book.Author
		b.Rating = book.Rating
		db.Create(&b)
		c.Status(201).JSON(fiber.Map{"message": "book successfully created"})
	}
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var payload *model.Book
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	if err := c.BodyParser(payload); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": fiber.ErrInternalServerError.Message})
	}
	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["Title"] = payload.Title
	}
	if payload.Author != "" {
		updates["Author"] = payload.Author
	}
	if payload.Rating != 0 {
		updates["Rating"] = payload.Rating
	}

	db.Update(updates).Where("ID", id)
	c.Send("Book Successfully updated")
}
