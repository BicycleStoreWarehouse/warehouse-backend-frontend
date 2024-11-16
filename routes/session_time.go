package routes

import (
	"log"
	"net/http"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveTime(c *gin.Context) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Baza danych nie została znaleziona w kontekście"})
		return
	}
	database := db.(*gorm.DB)

	var jsonData struct {
		UserID      int       `json:"user_id"`
		WorkedHours int       `json:"worked_hours"` // sekundy pracy
		BreakTime   int       `json:"break_time"`   // sekundy przerwy
		StartTime   time.Time `json:"start_time"`   // czas rozpoczęcia pracy
		EndTime     time.Time `json:"end_time"`     // czas zakończenia pracy
	}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		log.Printf("Błąd wiązania JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane wejściowe"})
		return
	}

	// Konwersja na time.Duration
	workedHours := time.Duration(jsonData.WorkedHours) * time.Second
	breakTime := time.Duration(jsonData.BreakTime) * time.Second

	log.Printf("Dane JSON po przetworzeniu: UserID: %d, WorkedHours: %d sekund, BreakTime: %d sekund, StartTime: %v, EndTime: %v",
		jsonData.UserID, jsonData.WorkedHours, jsonData.BreakTime, jsonData.StartTime, jsonData.EndTime)

	err := models.DailyTimekeeping(jsonData.UserID, workedHours, breakTime, jsonData.StartTime, jsonData.EndTime, database)
	if err != nil {
		log.Printf("Błąd zapisywania czasu pracy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas zapisywania czasu pracy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Czas pracy zapisany pomyślnie"})
}
