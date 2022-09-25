package database

import (
	"fmt"
	"log"
	"middleware/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/simple_api?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}
