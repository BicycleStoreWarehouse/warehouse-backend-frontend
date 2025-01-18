package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID            uint      `gorm:"not null"`          // ID pracownika, któremu przypisano zadanie
	User              User      `gorm:"foreignKey:UserID"` // Relacja z modelem User
	Description       string    `gorm:"not null"`          // Opis zadania
	Deadline          time.Time `gorm:"not null"`          // Data realizacji zadania
	Priority          string    `gorm:"not null"`          // Priorytet zadania (np. 1 = wysoki, 2 = średni, 3 = niski)
	IsCompleted       bool      `gorm:"default:false"`     // Status zadania (false = niewykonane, true = wykonane)
	DaysUntilDeadline int       `gorm:"-"`                 // Pole obliczane dynamicznie
}

// Tworzenie nowego zadania
func CreateTask(db *gorm.DB, userID uint, description string, deadline time.Time, priority string) (Task, error) {
	task := Task{
		UserID:      userID,
		Description: description,
		Deadline:    deadline,
		Priority:    priority,
		IsCompleted: false,
	}

	result := db.Create(&task)
	return task, result.Error
}

// Pobieranie niewykonanych zadań przypisanych do pracownika
func GetTasksByUserID(db *gorm.DB, userID uint) ([]Task, error) {
	var tasks []Task
	err := db.Where("user_id = ?", userID).Where("is_completed = ?", false).Order("deadline ASC, priority ASC").Find(&tasks).Error
	return tasks, err
}

// Oznaczanie zadania jako wykonane
func MarkTaskAsCompleted(db *gorm.DB, taskID uint, UserID uint) error {
	return db.Model(&Task{}).Where("id = ?", taskID).Update("is_completed", true).Error
}

// Pobieranie niewykonanych zadań pracowników
func GetUncompletedTasksForAllUsers(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Preload("User").Where("is_completed = ?", false).Find(&tasks).Error
	return tasks, err
}

// Pobieranie liczby niewykonanych zadań dla użytkownika
func GetUncompletedTasksCountByUserID(db *gorm.DB, userID uint) (int64, error) {
	var count int64
	err := db.Model(&Task{}).Where("user_id = ? AND is_completed = ?", userID, false).Count(&count).Error
	return count, err
}

func GetOverdueTasksCountByUserID(db *gorm.DB, userID uint) (int, error) {
	var count int64
	err := db.Model(&Task{}).
		Where("user_id = ? AND is_completed = ? AND deadline < ?", userID, false, time.Now()).
		Count(&count).Error
	return int(count), err
}

func CountIncompleteTasks(db *gorm.DB, dateFrom, dateTo *time.Time) (int64, error) {
	var incompleteTaskCount int64

	// Zbuduj zapytanie z opcjonalnym filtrowaniem dat
	query := db.Model(&Task{}).Where("is_completed = ?", false)
	if dateFrom != nil && dateTo != nil {
		query = query.Where("deadline BETWEEN ? AND ?", *dateFrom, *dateTo)
	} else if dateFrom != nil {
		query = query.Where("deadline >= ?", *dateFrom)
	} else if dateTo != nil {
		query = query.Where("deadline <= ?", *dateTo)
	}

	// Zlicz niewykonane zadania
	if err := query.Count(&incompleteTaskCount).Error; err != nil {
		return 0, err
	}

	return incompleteTaskCount, nil
}
