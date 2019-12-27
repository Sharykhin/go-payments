package gorm

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	G *gorm.DB
	g *gorm.DB
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

// NewGORMConnection returns a GORM connection
func NewGORMConnection() *gorm.DB {
	sync.Once{}.Do(func() {
		g = connect(true)
	})

	return g
}

// TODO: think about connect or init. Personally I would prefer init.
func connect(enableLogMode bool) *gorm.DB {
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

	db.LogMode(enableLogMode)

	return db
}
