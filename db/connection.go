package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=lithiumxp password=dbseg123 dbname=traductorapi port=5432"
var DB *gorm.DB
var err error

func DBConnection() {

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connected")
	}
}
