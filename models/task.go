package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`          // ID pracownika, któremu przypisano zadanie
	User        User      `gorm:"foreignKey:UserID"` // Relacja z modelem User
	Description string    `gorm:"not null"`          // Opis zadania
	Deadline    time.Time `gorm:"not null"`          // Data realizacji zadania
	Priority    string    `gorm:"not null"`          // Priorytet zadania (np. 1 = wysoki, 2 = średni, 3 = niski)
	IsCompleted bool      `gorm:"default:false"`     // Status zadania (false = niewykonane, true = wykonane)
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

// Pobieranie zadań przypisanych do pracownika
func GetTasksByUserID(db *gorm.DB, userID uint) ([]Task, error) {
	var tasks []Task
	err := db.Where("user_id = ?", userID).Order("deadline ASC, priority ASC").Find(&tasks).Error
	return tasks, err
}

// Oznaczanie zadania jako wykonane
func MarkTaskAsCompleted(db *gorm.DB, taskID uint, UserID uint) error {
	return db.Model(&Task{}).Where("id = ?", taskID).Update("is_completed", true).Error
}

func GetUncompletedTasksForAllUsers(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Preload("User").Where("is_completed = ?", false).Find(&tasks).Error
	return tasks, err
}
