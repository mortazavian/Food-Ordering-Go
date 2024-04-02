package repository

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB
var once sync.Once

func NewGormPostgres() *gorm.DB {
	once.Do(func() {
		// Load environment variables from .env file
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		// Get database connection information from environment variables
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		tehranTimezone, _ := time.LoadLocation("Asia/Tehran")

		// Connection configuration
		dsn := &url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword(dbUser, dbPassword),
			Host:     dbHost,
			Path:     dbName,
			RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
		}

		// Convert URL to connection string
		connStr := dsn.String()

		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		fmt.Println("Successfully connected to the database!")

		instance = db
	})

	return instance
}
