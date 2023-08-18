package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct { //defined Repository datatype
	DB *gorm.DB // declared member DB which is a pointer to gorm.DB datatype , and this gives ability to interact with database
}

func (r *Repository) CreateBook(context *fiber.Ctx) error { // created method for Repository struct , this method returns an error
	book := Book{}

	context.BodyParser(&book)
}

func (r *Repository) SetupRoutes(app *fiber.App) { // the datatype of value which we have passed as parameter to this function was initialized with value
	//fiber.New() which is of type *fiber.App i.e a pointer to value of datatype fiber.App
	// since instead of doing var a int = 4 we usually do a:= 4 , similar was done here
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBooksByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{ //defined r of datatype Repository
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}
