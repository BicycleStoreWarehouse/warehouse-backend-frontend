package controllers

import (
	"fmt"
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)

	} else if c.Request.Method == http.MethodPost {
		email := c.PostForm("email")
		password := c.PostForm("password")

		user_exists, err := models.UserExists(db, email)

		if !user_exists || err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Error": "Email or password is incorrect",
			})
		}

		user_password, user_name, user_position, err := models.GetInfoByEmail(db, email)

		if password != user_password || err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Error": "Email or password is incorrect",
			})
		}

		message := fmt.Sprintf("Witaj, %v, jesteś %v", user_name, user_position)

		// session.Set("user", email)
		// session.Save()

		c.HTML(http.StatusOK, "home.html", gin.H{"message": message})

	}
}
