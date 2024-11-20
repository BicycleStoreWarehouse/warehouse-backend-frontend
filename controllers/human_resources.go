package controllers

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HrDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_hr.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard_hr.html", gin.H{
		"user_name": user.Name,
	})
}

func ListWorkers(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "List Workers"})
}

func CreateWorker(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Create Worker"})
}

func DetailWorker(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Detail Worker"})
}

func ListOrdersHR(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "List Orders"})
}

func CreateOrder(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Create Order"})
}

func ListAplicationsHR(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "List Applications"})
}

func DetailApplication(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Detail Application"})
}

func UpdateApplication(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Update Application"})
}

func ShowReport(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"message": "Show Report"})
}