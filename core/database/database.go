package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	G *gorm.DB
)

func init() {
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
	))
	if err != nil {
		log.Panicf("failed to open a database connection: %v", err)
	}
	if err := db.DB().Ping(); err != nil {
		log.Panicf("failed to ping a database: %v", err)
	}

	db.LogMode(true)
	G = db

}
