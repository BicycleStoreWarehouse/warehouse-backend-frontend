package database

import (
	"log"
	"time"
	"warehouse/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadExampleData(db *gorm.DB) {
	twoYearsAgo := time.Now().AddDate(-2, 0, 0)
	threeYearsAgo := time.Now().AddDate(-3, 0, 0)

	models.CreateUser(db, "Jan", "Kowalski", "jan.kowalski@example.com", "Magazynowy", twoYearsAgo, "123456789", "hashed_password1")
	models.CreateUser(db, "Anna", "Nowak", "anna.nowak@example.com", "Magazynowy", twoYearsAgo, "987654321", "hashed_password2")
	models.CreateUser(db, "Piotr", "Wiśniewski", "piotr.wisniewski@example.com", "Magazynowy", twoYearsAgo, "555555555", "hashed_password3")

	models.CreateUser(db, "Kasia", "Wójcik", "kasia.wojcik@example.com", "HR", threeYearsAgo, "444444444", "hashed_password4")
	models.CreateUser(db, "Tomasz", "Kamiński", "tomasz.kaminski@example.com", "HR", threeYearsAgo, "333333333", "hashed_password5")
	models.CreateUser(db, "Ewa", "Zielińska", "ewa.zielinska@example.com", "Magazynowy", threeYearsAgo, "222222222", "hashed_password6")
}
func InitDatabase(dsn string) *gorm.DB {
	// Połączenie z bazą danych
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}

	// Automatyczne migracje
	if err := db.AutoMigrate(&models.WorkingHoursDaily{}, &models.User{}); err != nil {
		log.Fatalf("Błąd podczas migracji: %v", err)
	}

	// Dodatkowe migracje (np. zmiana kolumny)
	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	// Zmiana typu kolumny WorkedHours
	err := db.Exec("ALTER TABLE working_hours_dailies ALTER COLUMN worked_hours TYPE VARCHAR(8)").Error
	if err != nil {
		log.Fatalf("Nie udało się zmienić typu kolumny WorkedHours: %v", err)
	}

	// Zmiana typu kolumny BreakTime
	err = db.Exec("ALTER TABLE working_hours_dailies ALTER COLUMN break_time TYPE VARCHAR(8)").Error
	if err != nil {
		log.Fatalf("Nie udało się zmienić typu kolumny BreakTime: %v", err)
	}

	log.Println("Migracje zakończone pomyślnie.")
}
