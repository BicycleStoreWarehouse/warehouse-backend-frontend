package models

import (
	"time"

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
	Read            string `gorm:"default:'Nieodczytane'"` // 'Nieodczytane', 'Odczytane', lub null
}

func CreateVacation(db *gorm.DB, userID uint, dateFrom, dateTo string, dateCount int, status string) (Vacation, error) {
	vacation := Vacation{
		UserID:    userID,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		DateCount: dateCount,
		Status:    status,
		Read:      "Nieodczytane", // Puste podczas tworzenia nowego urlopu
	}

	result := db.Create(&vacation)
	return vacation, result.Error
}

func MarkVacationAsRead(db *gorm.DB, vacationID uint) error {
	return db.Model(&Vacation{}).Where("id = ?", vacationID).Update("read", "Odczytane").Error
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
	updateData := map[string]interface{}{
		"status": status,
		"read":   "Nieodczytane", // Automatycznie ustaw na "Nieodczytane" przy zmianie statusu
	}
	if rejectionReason != nil {
		updateData["rejection_reason"] = *rejectionReason
	}
	return db.Model(&Vacation{}).Where("id = ?", id).Updates(updateData).Error
}

// Pobieranie liczby niewykonanych zadań dla użytkownika
func GetVacationCountByUserID(db *gorm.DB, userID uint) (int64, error) {
	var count int64
	err := db.Model(&Vacation{}).Where("user_id = ? AND read = ?", userID, "Nieodczytane").Count(&count).Error
	return count, err
}

func CountUsersOnVacation(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Vacation{}).
		Where("date_from <= ? AND date_to >= ?", time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
