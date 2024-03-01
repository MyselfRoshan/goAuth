package database

import (
	"log"
	"os"

	"github.com/MyselfRoshan/goAuth/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URI")
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Couldnot connect to the database")
	}
	DB = con

	con.AutoMigrate(&models.User{})
}
