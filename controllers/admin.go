package controllers

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"user_name": user.Name,
	})
}
