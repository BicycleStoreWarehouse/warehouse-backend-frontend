package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkingHoursDaily struct {
	gorm.Model
	UserID      int  			`gorm:"not null"`
	User        User 			`gorm:"foreignKey:UserID"`
	Day 		string			`gorm:"not null"`
	WorkedHours time.Duration 	`gorm:"not null"`
}

type WorkingHoursMonthly struct {
	gorm.Model
	UserID      	 int  			`gorm:"not null"`
	User        	 User 			`gorm:"foreignKey:UserID"`
	Month			 string 		`gorm:"unique;not null"`
	TotalWorkedHours time.Duration	`gorm:"not null"`
}


func DailyTimekeeping(user_id int, time_duration time.Duration, db *gorm.DB) error {
	today := time.Now().Format("2006-01-02")

	daily_time := WorkingHoursDaily{
		UserID:			user_id,
		Day: 			today,
		WorkedHours:	time_duration,
	}

	result := db.Create(&daily_time)

	return result.Error
}


func MonthlyTimekeeping(user_id int, time_duration time.Duration, db *gorm.DB) error {
	var monthlyHours WorkingHoursMonthly
	month := time.Now().Format("2006-01")

	result := db.Where("user_id = ? AND month = ?", user_id, month).First(&monthlyHours)

	if result.Error == gorm.ErrRecordNotFound{
		monthly_timekeeping := WorkingHoursMonthly{
			UserID:				user_id,
			Month:				month,
			TotalWorkedHours:	time_duration,
		}

		db.Create(&monthly_timekeeping)
	} else {
		monthlyHours.TotalWorkedHours += time_duration
		result = db.Save(&monthlyHours)
	}
	return result.Error
}