package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"

	// "net/http"
	"log"
	// "errors"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	store, err := NewPostgresStore(config)
	if err != nil {
		log.Fatal(err)
	}

	store.CreateArticleTable()
	if err != nil {
		log.Fatal(err)
	}

	// clothes := []Clothing{}

	PORT := os.Getenv("PORT")

	app := fiber.New()

	// Get All Clothes
	app.Get("/api/clothes", getArticleHandler(store))

	// Create Article
	app.Post("/api/clothes", createArticleHandler(store))

	// // Wear Article
	// router.Patch("/api/clothes/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")
	// 	for i, article := range clothes {
	// 		if fmt.Sprint(article.ID) == id {
	// 			clothes[i].IncrementWears()
	// 			return c.Status(200).JSON(clothes[i])
	// 		}
	// 	}
	// 	return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	// })

	// // Delete an Article
	// router.Delete("/api/clothes/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")
	// 	for i, article := range clothes {
	// 		if fmt.Sprint(article.ID) == id {
	// 			clothes = append(clothes[:i], clothes[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{"sucess": true})
	// 		}
	// 	}
	// 	return c.Status(404).JSON(fiber.Map{"error": "Article not found"})

	// })

	log.Fatal(app.Listen(":" + PORT))
}

// Post /article
func createArticleHandler(store Storage) fiber.Handler {
	return func(c *fiber.Ctx) error {
		article := new(Clothing)
		if err := c.BodyParser(article); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		article.UpdateCPW() // Update the cost-per-wear if this is a part of your logic

		if err := store.CreateArticle(article); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create article"})
		}

		return c.Status(fiber.StatusCreated).JSON(article)
	}
}

// Get /article
func getArticleHandler(store Storage) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clothes, err := store.GetClothing()
		if err != nil {
			return err
		}
		return c.Status(200).JSON(clothes)
	}
}
