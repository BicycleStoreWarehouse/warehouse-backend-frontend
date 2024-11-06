package controllers

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user_password, err := models.GetUserPassword(db, email)

	if password != user_password || err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Email lub hasło jest nieprawidłowe",
		})
		return
	}

	user_position, _ := models.GetUserPosition(db, email)

	session := sessions.Default(c)
	session.Set("user_email", email)
	session.Save()

	if user_position == "Magazynowy" {
		c.Redirect(http.StatusFound, "/warehouse/dashboard")
	} else if user_position == "HR" {
		c.Redirect(http.StatusFound, "/hr/dashboard")
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Twoje stanowisko jest nieprawidłowe",
		})
		return
	}
}
