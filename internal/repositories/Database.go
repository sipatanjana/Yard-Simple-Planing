package repositories

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"YardPlaning/internal/entities"
)

var errorNoData = errors.New("no data")

func InitDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = CreateTable(db) //untuk inisialisasi tabel

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func CreateTable(db *gorm.DB) error {
	db.AutoMigrate(
		&entities.Yard{},
		&entities.Block{},
		&entities.YardPlan{},
	)
	return nil

	//data migrastion untuk user

}
