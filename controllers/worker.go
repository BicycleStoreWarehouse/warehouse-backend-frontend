package controllers

import (
	"log"
	"net/http"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WorkerDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_warehouse.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard_warehouse.html", gin.H{
		"user_name": user.Name,
	})
}

func DashboardWorker(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_worker.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard_worker.html", gin.H{
		"user_name":        user.Name,
		"user_surname":     user.Surname,
		"user_email":       user.Email,
		"user_phone":       user.Phone,
		"user_street":      user.Street,
		"user_city":        user.City,
		"user_state":       user.State,
		"user_zip":         user.Zip,
		"user_country":     user.Country,
		"user_bankAccount": user.BankAccount,
		"user_nameBank":    user.NameBank,
	})
}

func SaveTime(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	var workAndBreakTime struct {
		WorkedHours int `json:"worked_hours"`
		BreakTime   int `json:"break_time"`
	}

	if err := c.ShouldBindJSON(&workAndBreakTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane wejściowe"})
		return
	}

	workedHours := time.Duration(workAndBreakTime.WorkedHours) * time.Second
	breakTime := time.Duration(workAndBreakTime.BreakTime) * time.Second

	err := models.DailyTimekeeping(userID.(uint), workedHours, breakTime, db)

	if err != nil {
		log.Printf("Błąd zapisywania czasu pracy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas zapisywania czasu pracy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Czas pracy zapisany pomyślnie"})
}

func TimeTracking(c *gin.Context, db *gorm.DB) {
	userID, exists_id := c.Get("user_id")

	if !exists_id {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Musisz być zalogowany",
		})
		return
	}

	daily_report, err := models.GetDailyReportForUser(db, userID.(uint))

	if err != nil {
		c.HTML(http.StatusUnauthorized, "time_tracking.html", gin.H{
			"message": "Błąd połączenia z bazą danych",
		})
		return
	}

	c.HTML(http.StatusOK, "time_tracking.html", gin.H{
		"WorkHistory": daily_report,
	})
}

func ListProducts(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_products.html", gin.H{"message": "List Products"})
}

func CreateProduct(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "create_product.html", gin.H{"message": "Create Product"})
}

func ListOrdersWorker(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_orders_worker.html", gin.H{"message": "List Orders"})
}

func DetailOrder(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "detail_order.html", gin.H{"message": "Detail Order"})
}

func UpdateOrder(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "update_order.html", gin.H{"message": "Update Order"})
}

func ListAplicationsWorker(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_applications_worker.html", gin.H{"message": "List Applications"})
}

func CreateApplication(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "create_application.html", gin.H{"message": "Create Application"})
}
