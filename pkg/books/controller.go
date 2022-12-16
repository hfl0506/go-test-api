package books

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/hfl0506/go-test-api/pkg/common/models"
	"gorm.io/gorm"
	"github.com/hfl0506/go-test-api/pkg/common/"
)
type handler struct {
	DB *gorm.DB
}
	
type AddBookRequestBody struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
	Rating int `json:"rating"`
}

func RegisterRoutes(app fiber.Router, db *gorm.DB) {
	h := &handler {
		DB: db,
	}

	routes := app.Group("/books")

	routes.Post("/", h.AddBook)
	routes.Get("/", h.GetBooks)
	routes.Get("/:id", h.GetBook)
	routes.Put("/:id", h.UpdateBook)
	routes.Post("/:id", h.DeleteBook)

}

// get all
func (h *handler) GetBooks(c *fiber.Ctx) error {
	var books []*models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}

// get one
func (h *handler) GetBook(c *fiber.Ctx) error {
	id := 
	var book *models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	
	return c.Status(fiber.StatusOK).JSON(&book)
}

//create
func (h *handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	fmt.Println("before parser body")

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fmt.Println("after parser body")

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.Rating = body.Rating

	fmt.Println(book)

	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	fmt.Println("after create book")

	return c.Status(fiber.StatusCreated).JSON(&book)
}

// update one

func (h *handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	body := AddBookRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book *models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Title = body.Title
	book.Description = body.Description
	book.Author = body.Author
	book.Rating = body.Rating

	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}

// delete one
func (h *handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	
	var book *models.Book

	if result := h.DB.First(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&book)

	return c.Status(fiber.StatusOK).SendString("book"+id+"has been delete")
}