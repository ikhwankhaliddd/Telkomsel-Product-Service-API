package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	if err := runMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}

func runMigrations(db *sqlx.DB) error {
	migrationSQL := `
	CREATE TABLE categories (
		id serial PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP;
	);
	
	CREATE TABLE products (
		id serial PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		rating INT,
		category_id INT REFERENCES categories(id),
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP;
	);
	
	CREATE TABLE variety (
		id serial PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price FLOAT,
		stock INT,
		image VARCHAR(255),
		product_id INT REFERENCES products(id),
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP;
	);
	`

	_, err := db.Exec(migrationSQL)
	if err != nil {
		return err
	}

	return nil
}
