package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type WorkingHoursDaily struct {
	gorm.Model
	UserID      int       `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	Day         string    `gorm:"not null"`
	WorkedHours string    `gorm:"type:varchar(8);not null"`
	BreakTime   string    `gorm:"type:varchar(8);not null"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

type TimeTracking struct {
	gorm.Model
	UserID      int `gorm:"not null"`
	WorkingTime int
	BreakTime   int
}

type WorkingHoursMonthly struct {
	gorm.Model
	UserID           int    `gorm:"not null"`
	User             User   `gorm:"foreignKey:UserID"`
	Month            string `gorm:"unique;not null"`
	TotalWorkedHours string `gorm:"type:varchar(8);not null"`
}

func DailyTimekeeping(user_id int, worked_hours time.Duration, break_time time.Duration, startTime time.Time, endTime time.Time, db *gorm.DB) error {
	today := time.Now().Format("2006-01-02")

	// Konwersja czasu na format hh:mm:ss
	workedHoursStr := formatDurationToHHMMSS(worked_hours)
	breakTimeStr := formatDurationToHHMMSS(break_time)

	daily_time := WorkingHoursDaily{
		UserID:      user_id,
		Day:         today,
		WorkedHours: workedHoursStr,
		BreakTime:   breakTimeStr,
		StartTime:   startTime,
		EndTime:     endTime,
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

// Funkcja pomocnicza do konwersji time.Duration na hh:mm:ss
func formatDurationToHHMMSS(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
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
