package db

import (
	"fmt"
	"go-mt2/pkg/db/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBService struct {
	AuthDB *gorm.DB
	GameDB *gorm.DB
}

func NewAuthDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_ROOT_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("AUTH_DB_DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewGameDB() *gorm.DB {
	godotenv.Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_ROOT_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("AUTH_DB_DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DBService) MigrateAuth() {
	db.AuthDB.AutoMigrate(&models.Account{})
	db.AuthDB.AutoMigrate(&models.AccountStatus{})

	db.AuthDB.Create(&models.Account{
		Username:        "admin",
		Password:        "admin",
		Email:           "admin@example.com",
		AccountStatusId: 1,
	})
}

func (db *DBService) MigrateGame() {
	db.AuthDB.Raw()
}

func (db *DBService) SeedAuth() {
	db.AuthDB.Create(&models.AccountStatus{
		Id:           1,
		ClientStatus: "OK",
		AllowLogin:   true,
		Description:  "Default Status",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	db.AuthDB.Create(&models.Account{
		Id:              1,
		Username:        "admin",
		Password:        "admin",
		Email:           "admin@example.com",
		DeleteCode:      "1234567",
		AccountStatusId: 1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

}
