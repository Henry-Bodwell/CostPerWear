package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateArticle(*Clothing) error
	DeleteArticle(int) error
	UpdateArticle(*Clothing) error
	GetArticleByID(int) (*Clothing, error)
	GetClothing() ([]*Clothing, error)
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(config Config) (*PostgresStore, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	// Open the server see if there is an issue
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the server, see if its connectable
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)

	fmt.Printf("Listing on port: %s...", config.Port)
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return nil
}

func (s *PostgresStore) CreateArticleTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS WebApp.clothing (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			price NUMERIC(10, 2),
			wears SMALLINT DEFAULT 0,
			material VARCHAR(50),
			brand VARCHAR(50),
			season VARCHAR(20),
			costPerWear DOUBLE PRECISION,
			clothingType VARCHAR(50),
			image TEXT,
			lastWorn TIMESTAMP
		)
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateArticle(c *Clothing) error {
	query := `INSERT INTO WebApp.clothing
		(name, price, wears, material, brand, season, costPerWear, clothingType, image, lastWorn)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`
	resp, err := s.db.Query(
		query,
		c.Name,
		c.Price,
		c.Wears,
		c.Material,
		c.Brand,
		c.Season,
		c.CostPerWear,
		c.ClothingType,
		c.Image,
		c.LastWorn)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) DeleteArticle(int) error {
	return nil
}

func (s *PostgresStore) UpdateArticle(*Clothing) error {
	return nil
}

func (s *PostgresStore) GetArticleByID(int) (*Clothing, error) {
	return nil, nil
}

func (s *PostgresStore) GetClothing() ([]*Clothing, error) {
	rows, err := s.db.Query("SELECT * FROM WebApp.clothing")
	if err != nil {
		return nil, err
	}

	articles := []*Clothing{}
	for rows.Next() {
		article := new(Clothing)
		if err := rows.Scan(
			&article.ID,
			&article.Name,
			&article.Price,
			&article.Wears,
			&article.Material,
			&article.Brand,
			&article.Season,
			&article.CostPerWear,
			&article.ClothingType,
			&article.Image,
			&article.LastWorn,
		); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
