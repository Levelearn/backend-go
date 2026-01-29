package config

import (
	"fmt"
	"levelearn-backend/internal/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Ambil data dari env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT") // Tambahan Port
	dbName := os.Getenv("DB_NAME")

	// Format DSN (Data Source Name) MySQL
	// referensi: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Panic jika database tidak konek, karena backend tidak bisa jalan tanpa DB
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto Migrate (Sama seperti sebelumnya)
	log.Println("Database connected successfully, running migrations...")
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Course{},
		&entity.UserCourse{},
		&entity.Chapter{},
		&entity.Material{},
		&entity.Assessment{},
		&entity.Assignment{},
		&entity.UserChapter{},
		&entity.Badge{},
		&entity.UserBadge{},
		&entity.Trade{},
		&entity.UserTrade{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}