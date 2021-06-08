package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	//db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	db, err = gorm.Open(sqlite.Open("/tmp/gorm.db"), &gorm.Config{})
	if err != nil {
		log.Printf("%+v", err)
	}
	//db.LogMode(true)

	db.AutoMigrate(&MQTTConf{}) // TODO remember to add each model here
	db.AutoMigrate(&TiltConf{})
	db.AutoMigrate(&Consumers{})
}

func GetDB() *gorm.DB {
	return db
}
