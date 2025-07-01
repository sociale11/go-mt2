package db

import (
	"fmt"
	"go-mt2/pkg/db/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBService struct {
	AuthDB *gorm.DB
	GameDB *gorm.DB
}

func Init() {

}

func NewAuthDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_ROOT_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("AUTH_DB_DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return db
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
		Password:        "$2b$05$KXeREc2TNuUR6IcgzUiX4.WA/0i3Yd3WpUHMtAcQi1ojWRdeQ9ExS",
		Email:           "admin@example.com",
		DeleteCode:      "1234567",
		LastLogin:       time.Now(),
		AccountStatusId: 1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

}
