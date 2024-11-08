package controllers

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HrDashboard(c *gin.Context, db *gorm.DB) {
	user_email, exists := c.Get("user_email")

	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user_name, _, err := models.GetUserNameAndSurname(db, user_email.(string))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_hr.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard_hr.html", gin.H{
		"user_name": user_name,
	})
}