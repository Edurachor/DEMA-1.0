package database

import (
	"apidema/api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	USER   = "postgres"
	PASS   = "root"
	HOST   = "localhost"
	PORT   = 5432
	DBNAME = "dema"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", HOST, USER, PASS, DBNAME, PORT)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados!")
}

func AutoMigration() {
	Db.Migrator().DropTable(&models.User{})
	Db.Migrator().AutoMigrate(&models.User{})
}
