package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	// Skonfiguruj DSN z użyciem hosta, aby upewnić się, że połączenie będzie poprawne
	dsn := "host=db user=group15 password=group15 dbname=group15 port=5432 sslmode=disable TimeZone=Europe/Warsaw"

	// Otwórz połączenie z bazą danych
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Błąd połączenia z bazą danych: %v", err)
		return nil, err // Upewnij się, że zwracasz nil, jeśli wystąpił błąd
	}

	return db, nil // Zwróć db oraz nil jako błąd
}
