package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func InitDB() *gorm.DB {
	DBURL := "postgresql://postgres:abcde12345@localhost:5432/go-ecommerce"
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return db
}
