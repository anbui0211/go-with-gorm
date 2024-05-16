package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() (conn *gorm.DB, err error) {
	dsn := "user:pass@tcp(127.0.0.1:14306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	log.Println("Connecting to database successfully ...")
	return
}
