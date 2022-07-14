package db

import (
	"GoAPI/config"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err)
		return nil, err
	}
	if config.DBDriver == "postgres" {
		return gorm.Open(postgres.Open(config.DBSource))
	}

	log.Fatal("Driver not allowed! ", config.DBDriver)
	return nil, errors.New("Driver not allowed!")
}
