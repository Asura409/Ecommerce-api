// pkg/database/database.go
package database

import (
	
	"log"
	"os"
	"ecommerce-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection and runs migrations
func Connect() {
	var err error
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=ecommerce port=5432 sslmode=disable TimeZone=UTC"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run migrations
	if err := AutoMigrate(DB); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
}




// AutoMigrate runs all database migrations
func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}