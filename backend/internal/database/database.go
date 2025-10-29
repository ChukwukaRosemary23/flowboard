package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ChukwukaRosemary23/flowboard-backend/config"
	"github.com/ChukwukaRosemary23/flowboard-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase connects to PostgreSQL and runs migrations
func ConnectDatabase(cfg *config.Config) error {
	var dsn string

	// Check if DATABASE_URL exists (Render uses this)
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL != "" {
		// Use DATABASE_URL (for Render/production)
		dsn = databaseURL
		log.Println("📦 Using DATABASE_URL for connection")
	} else {
		// Use individual variables (for local development)
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
		)
		log.Println("🏠 Using individual DB variables for connection")
	}

	// Set up GORM config
	gormConfig := &gorm.Config{}

	// In development, show SQL queries (helpful for learning!)
	if cfg.Env == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// Connect to database
	database, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Store in global variable
	DB = database
	log.Println("✅ Database connected successfully!")

	// Run migrations (create tables)
	if err := runMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// runMigrations creates/updates all database tables
func runMigrations() error {
	log.Println("🔄 Running database migrations...")

	// AutoMigrate creates tables if they don't exist
	// It also updates table structure if models changed
	err := DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Card{},
		&models.Label{},
		&models.Comment{},
		&models.Attachment{},
		&models.CardMember{},
		&models.CardLabel{},
		&models.Activity{},
	)

	if err != nil {
		return err
	}

	log.Println("✅ Database migrations completed successfully!")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
