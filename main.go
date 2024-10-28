package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"path/filepath"
	"runtime"

	// "net/http"
	"log"

	"github.com/gofiber/fiber/v2"
	// "errors"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func newDB(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	// Open the server see if there is an issue
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the server, see if its connectable
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	return db, nil
}

func main() {

	_, filename, _, _ := runtime.Caller(0)
	// Navigate up two levels (from cmd/myapp to root)
	projectRoot := filepath.Join(filepath.Dir(filename), "../..")

	// Load the .env file from project root
	err := godotenv.Load(filepath.Join(projectRoot, ".env"))
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

	db, err := newDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	clothes := []Clothing{}

	PORT := os.Getenv("PORT")

	router := fiber.New()

	router.Get("/api/clothes", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(clothes)
	})

	// Create Article
	router.Post("/api/clothes", func(c *fiber.Ctx) error {
		article := &Clothing{}

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
