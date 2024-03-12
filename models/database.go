package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgres://postgress:123@localhost:5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	DB = db
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}
}
