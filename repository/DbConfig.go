package repository

import (
	"fmt"
	"os"

	"github.com/MicroSOA-09/blog-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitDB() *gorm.DB {

	server := getEnv("DATABASE_HOST", "localhost")
	port := getEnv("DATABASE_PORT", "5432")
	database := getEnv("DATABASE_SCHEMA", "explorer-v1")
	schema := getEnv("DATABASE_SCHEMA_NAME", "blog")
	user := getEnv("DATABASE_USERNAME", "postgres")
	password := getEnv("DATABASE_PASSWORD", "super")
	integratedSecurity := getEnv("DATABASE_INTEGRATED_SECURITY", "false")
	sslmode := "disable"
	if integratedSecurity == "true" {
		sslmode = "require"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s search_path=%s user=%s password=%s sslmode=%s",
		server, port, database, schema, user, password, sslmode,
	)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("failed to connect to database" + error.Error())
	}

	db.AutoMigrate(&model.BlogPost{})
	return db

}
