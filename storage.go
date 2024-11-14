package main

import (
	"database/sql"
	"fmt"
	"log"

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

	log.Println("Listing on port:", config.Port)

	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	err := s.CreateArticleTable()
	if err != nil {
		return err
	}
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
			lastWorn TIMESTAMP,
			deleted BOOLEAN DEFAULT FALSE
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
	_, err := s.db.Exec(
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
		log.Printf("Error executing query: %v", err)
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteArticle(id int) error {
	_, err := s.db.Exec(`UPDATE WebApp.clothing
						SET deleted = TRUE
						WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) UpdateArticle(c *Clothing) error {
	query := `
        UPDATE WebApp.clothing 
        SET name = $1,
            price = $2,
            wears = $3,
            material = $4,
            brand = $5,
            season = $6,
            costPerWear = $7,
            clothingType = $8,
            image = $9,
            lastWorn = $10
        WHERE id = $11
    `

	result, err := s.db.Exec(
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
		c.LastWorn,
		c.ID)
	if err != nil {
		log.Printf("Error updating article: %v", err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("article %d not found", c.ID)
	}

	return nil
}

func (s *PostgresStore) GetArticleByID(id int) (*Clothing, error) {
	query := `SELECT * FROM WebApp.clothing WHERE id = $1 AND deleted = FALSE`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoArticle(rows)
	}

	return nil, fmt.Errorf("article %d not found", id)
}

func (s *PostgresStore) GetClothing() ([]*Clothing, error) {
	query := `SELECT * FROM WebApp.clothing WHERE deleted = FALSE ORDER BY lastWorn DESC`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*Clothing{}
	for rows.Next() {
		article, err := scanIntoArticle(rows)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func scanIntoArticle(rows *sql.Rows) (*Clothing, error) {
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
		&article.Deleted,
	); err != nil {
		return nil, err
	}

	return article, nil
}
