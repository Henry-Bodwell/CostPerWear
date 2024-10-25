package main

import (
	"fmt"
	"os"

	"github.com/Henry-Bodwell/CostPerWear/internal/app"
	"github.com/joho/godotenv"

	"path/filepath"
	"runtime"

	// "net/http"
	"log"

	"github.com/gofiber/fiber/v2"
	// "errors"
)

// func getArticles(C *gin.Context) {
// 	// Go to database get articles of clothes

// 	// C.IndentedJSON(http.StatusOK, clothes)
// }

func main() {

	_, filename, _, _ := runtime.Caller(0)
	// Navigate up two levels (from cmd/myapp to root)
	projectRoot := filepath.Join(filepath.Dir(filename), "../..")

	// Load the .env file from project root
	err := godotenv.Load(filepath.Join(projectRoot, ".env"))
	if err != nil {
		log.Fatal("Error loading .env")
	}

	router := fiber.New()

	clothes := []app.Clothing{}

	PORT := os.Getenv("PORT")

	router.Get("/api/clothes", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(clothes)
	})

	// Create Article
	router.Post("/api/clothes", func(c *fiber.Ctx) error {
		article := &app.Clothing{}

		if err := c.BodyParser(article); err != nil {
			return err
		}

		article.UpdateCPW()

		if article.Name == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
		}

		article.ID = len(clothes) + 1
		clothes = append(clothes, *article)

		return c.Status(201).JSON(article)

	})

	// Wear Article
	router.Patch("/api/clothes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, article := range clothes {
			if fmt.Sprint(article.ID) == id {
				clothes[i].IncrementWears()
				return c.Status(200).JSON(clothes[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	})

	// Delete an Article
	router.Delete("/api/clothes/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, article := range clothes {
			if fmt.Sprint(article.ID) == id {
				clothes = append(clothes[:i], clothes[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"sucess": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})

	})

	log.Fatal(router.Listen(":" + PORT))
}
