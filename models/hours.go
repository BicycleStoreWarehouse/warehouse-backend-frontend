package models

import (
	"log"
	"time"

	"gorm.io/gorm"

	"warehouse/helpers"
)

type WorkingHoursDaily struct {
	gorm.Model
	UserID      int    `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	Day         string `gorm:"not null"`
	WorkedHours string `gorm:"type:varchar(8);not null"`
	BreakTime   string `gorm:"type:varchar(8);not null"`
}

type WorkingHoursMonthly struct {
	gorm.Model
	UserID           int    `gorm:"not null"`
	User             User   `gorm:"foreignKey:UserID"`
	Month            string `gorm:"unique;not null"`
	TotalWorkedHours string `gorm:"type:varchar(8);not null"`
}

func DailyTimekeeping(user_id int, worked_hours time.Duration, break_time time.Duration, db *gorm.DB) error {
	today := time.Now().Format("2006-01-02")

	workedHoursStr := helpers.FormatDurationToHHMMSS(worked_hours)
	breakTimeStr := helpers.FormatDurationToHHMMSS(break_time)

	daily_time := WorkingHoursDaily{
		UserID:      user_id,
		Day:         today,
		WorkedHours: workedHoursStr,
		BreakTime:   breakTimeStr,
	}

	log.Printf("Zapisujemy czas pracy: UserID: %d, Dzień: %s, WorkedHours: %s, BreakTime: %s", user_id, today, workedHoursStr, breakTimeStr)
	result := db.Create(&daily_time)
	if result.Error != nil {
		log.Printf("Błąd podczas zapisywania rekordu czasu pracy: %v", result.Error)
		return result.Error
	}
	log.Println("Rekord czasu pracy zapisany pomyślnie")
	return nil
}

func MonthlyTimekeeping(user_id int, time_duration string, db *gorm.DB) error {
	var monthlyHours WorkingHoursMonthly
	month := time.Now().Format("2006-01")

	result := db.Where("user_id = ? AND month = ?", user_id, month).First(&monthlyHours)

	if result.Error == gorm.ErrRecordNotFound {
		monthly_timekeeping := WorkingHoursMonthly{
			UserID:           user_id,
			Month:            month,
			TotalWorkedHours: time_duration,
		}

		db.Create(&monthly_timekeeping)
	} else {
		monthlyHours.TotalWorkedHours += time_duration
		result = db.Save(&monthlyHours)
	}
	return result.Error
}

func GetDailyReportForUser(db *gorm.DB, userID int) ([]WorkingHoursDaily, error) {

	var workingHours []WorkingHoursDaily

	// Pobieranie wszystkich rekordów dla danego użytkownika
	result := db.Where("user_id = ?", userID).Find(&workingHours)
	if result.Error != nil {
		log.Printf("Błąd podczas pobierania danych godzin pracy dla użytkownika %d: %v", userID, result.Error)
		return nil, result.Error
	}

	// Zwracanie pobranych rekordów
	return workingHours, nil
}
