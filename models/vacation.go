package models

import (
	"gorm.io/gorm"
)

type Vacation struct {
	gorm.Model
	UserID          uint   `gorm:"not null"`
	User            User   `gorm:"foreignKey:UserID"`
	DateFrom        string `gorm:"not null"`
	DateTo          string `gorm:"not null"`
	DateCount       int    `gorm:"not null"`
	Status          string `gorm:"default:'Wysłany'"`
	RejectionReason string `gorm:"default:null"`
}

func CreateVacation(db *gorm.DB, userID uint, dateFrom, dateTo string, dateCount int, status string) (Vacation, error) {
	vacation := Vacation{
		UserID:    userID,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		DateCount: dateCount,
		Status:    status,
	}

	result := db.Create(&vacation)
	return vacation, result.Error
}

func GetVacationByID(db *gorm.DB, id uint) (Vacation, error) {
	var vacation Vacation
	err := db.Preload("User").Where("id = ?", id).First(&vacation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Vacation{}, gorm.ErrRecordNotFound
		}
		return Vacation{}, err
	}
	return vacation, nil
}

func GetVacationsByUserID(db *gorm.DB, userID uint) ([]Vacation, error) {
	var vacations []Vacation
	err := db.Where("user_id = ?", userID).Find(&vacations).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return vacations, nil
}

func GetAllVacations(db *gorm.DB) ([]Vacation, error) {
	var vacations []Vacation
	err := db.Where("status = ?", "Wysłany").Preload("User").Find(&vacations).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return vacations, nil
}

func UpdateVacationDates(db *gorm.DB, id uint, dateFrom, dateTo string, dateCount int) error {
	var vacation Vacation
	err := db.First(&vacation, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	vacation.DateFrom = dateFrom
	vacation.DateTo = dateTo
	vacation.DateCount = dateCount

	return db.Save(&vacation).Error
}

func DeleteVacation(db *gorm.DB, id uint) error {
	return db.Delete(&Vacation{}, id).Error
}

func UpdateVacationStatus(db *gorm.DB, id uint, status string, rejectionReason *string) error {
	updateData := map[string]interface{}{"status": status}
	if rejectionReason != nil {
		updateData["rejection_reason"] = *rejectionReason
	}
	return db.Model(&Vacation{}).Where("id = ?", id).Updates(updateData).Error
}
