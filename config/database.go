package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	return connection
}
