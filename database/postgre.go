package database

import (
	"MyGarm/helpers"
	"MyGarm/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	helpers.LoadENV()
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connected successfully")
	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
