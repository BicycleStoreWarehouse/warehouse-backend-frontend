package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type WorkingHoursDaily struct {
	gorm.Model
	UserID      int           `gorm:"not null"`
	User        User          `gorm:"foreignKey:UserID"`
	Day         string        `gorm:"not null"`
	WorkedHours time.Duration `gorm:"not null"`
	BreakTime   time.Duration `gorm:"not null"`
}

type TimeTracking struct {
	gorm.Model
	UserID      int `gorm:"not null"`
	WorkingTime int // w sekundach
	BreakTime   int // w sekundach
}

type WorkingHoursMonthly struct {
	gorm.Model
	UserID           int           `gorm:"not null"`
	User             User          `gorm:"foreignKey:UserID"`
	Month            string        `gorm:"unique;not null"`
	TotalWorkedHours time.Duration `gorm:"not null"`
}

func DailyTimekeeping(user_id int, worked_hours time.Duration, break_time time.Duration, db *gorm.DB) error {
	today := time.Now().Format("2006-01-02")

	daily_time := WorkingHoursDaily{
		UserID:      user_id,
		Day:         today,
		WorkedHours: worked_hours,
		BreakTime:   break_time,
	}

	log.Printf("Zapisujemy czas pracy: UserID: %d, Dzień: %s, WorkedHours: %v, BreakTime: %v", user_id, today, worked_hours, break_time)
	result := db.Create(&daily_time)
	if result.Error != nil {
		log.Printf("Błąd podczas zapisywania rekordu czasu pracy: %v", result.Error)
		return result.Error
	}
	log.Println("Rekord czasu pracy zapisany pomyślnie")
	return nil
}

func MonthlyTimekeeping(user_id int, time_duration time.Duration, db *gorm.DB) error {
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
