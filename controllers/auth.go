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

	user, err := models.GetUserByEmail(db, email)
	if err != nil || user.Password != password {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Email lub hasło jest nieprawidłowe",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID) // Ustawienie user_id w sesji
	session.Save()

	if user.Position == "Magazynowy" {
		c.Redirect(http.StatusFound, "/warehouse/dashboard")
	} else if user.Position == "HR" {
		c.Redirect(http.StatusFound, "/hr/dashboard")
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Twoje stanowisko jest nieprawidłowe",
		})
		return
	}
}
