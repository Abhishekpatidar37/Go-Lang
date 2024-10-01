package initializers

import (
	"fmt"
	"golang-crud/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func migrateToDb() {
	// Migrate the schema, including relationships
	migrationError := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Company{})
	if migrationError != nil {
		fmt.Println("Error in autoMigrate : ", migrationError)
	}
}

// To print database queries logs on terminal
func getLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Log SQL queries slower than this threshold
			LogLevel:                  logger.Info, // Log level (Info, Warn, Error)
			IgnoreRecordNotFoundError: true,        // Ignore record not found errors
			Colorful:                  true,        // Enable color logging
		},
	)
	return newLogger
}

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(), // Assign the logger
	})

	if err != nil {
		log.Println("failed to connect database!")
	}

	migrateToDb()

	// Retrieve and print the current database name
	var dbName string
	DB.Raw("SELECT current_database()").Scan(&dbName)
	fmt.Println("Connected to postgres db:", dbName)
}
