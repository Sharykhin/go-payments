package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func init() {
	conn, err := gorm.Open("mysql", "payments")
	if err != nil {
		log.Printf("failed to open a database connection: %v", err)
	}

	db = conn
}