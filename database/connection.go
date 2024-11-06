package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	dsn := "host=db user=group15 password=group15 dbname=group15 port=5432 sslmode=disable TimeZone=Europe/Warsaw"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Błąd połączenia z bazą danych: %v", err)
		return nil, err
	}

	return db, nil
}
