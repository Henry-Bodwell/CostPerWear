package main

import (
	"os"

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

	server := newAPIServer(store, ":"+PORT)
	server.Run()
}
