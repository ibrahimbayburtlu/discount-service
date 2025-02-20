package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// LoadEnv loads environment variables from a .env file (optional)
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file could not be loaded, using default environment variables.")
	}
}

// ConnectDB initializes the database connection
func ConnectDB() (*gorm.DB, error) {
	// Load environment variables first
	LoadEnv()

	// Get database URL from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=discountdb port=5432 sslmode=disable"
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
	})
	if err != nil {
		log.Fatalf("🚨 Failed to connect to database: %v", err)
		return nil, err
	}

	// Get the database instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("🚨 Failed to get DB instance: %v", err)
		return nil, err
	}

	// Test the connection
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("🚨 Database is not accessible: %v", err)
		return nil, err
	}

	// Set database connection settings
	sqlDB.SetMaxOpenConns(10)   // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)    // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Connection lifetime (0 = unlimited)

	log.Println("✅ Database connected successfully!")
	DB = db
	return db, nil
}
