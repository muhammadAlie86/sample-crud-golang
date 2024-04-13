package database

import (
	"crud-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	InitDatabase()
	GetDB() *gorm.DB
}

type database struct {
	DB *gorm.DB
}

func NewDatabase() IDatabase {

	dsn := "host=localhost user=postgres password=alie1992 dbname=green_cake port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return &database{
		DB: db,
	}
}

func (d *database) InitDatabase() {
	// Run migrations here
	d.DB.AutoMigrate(&models.User{})
}

func (d *database) GetDB() *gorm.DB {
	return d.DB
}
