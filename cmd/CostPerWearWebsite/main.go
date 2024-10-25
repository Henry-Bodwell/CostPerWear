package main

import (
	"fmt"

	"github.com/Henry-Bodwell/CostPerWear/internal/app"

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
	router := fiber.New()

	clothes := []app.Clothing{}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "OK "})
	})

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

	log.Fatal(router.Listen("localhost:9090"))
}
