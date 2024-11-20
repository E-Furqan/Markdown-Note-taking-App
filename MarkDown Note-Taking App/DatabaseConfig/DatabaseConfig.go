package databaseConfig

import (
	"fmt"
	"log"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	model.DatabaseEnvVar
}

func NewDatabase(env model.DatabaseEnvVar) *DatabaseConfig {
	return &DatabaseConfig{
		DatabaseEnvVar: env,
	}
}

var DB *gorm.DB

func (DatabaseConfig *DatabaseConfig) Connection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var connection_string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DatabaseConfig.DatabaseEnvVar.HOST, DatabaseConfig.DatabaseEnvVar.PORT,
		DatabaseConfig.DatabaseEnvVar.USER, DatabaseConfig.DatabaseEnvVar.PASSWORD,
		DatabaseConfig.DatabaseEnvVar.DB_NAME)

	DB, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&model.Notes{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established")

	return DB
}
