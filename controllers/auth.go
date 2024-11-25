package controllers

import (
	"log"
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

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User not found for email: %s", email)
		} else {
			log.Printf("Error fetching user for email %s: %v", email, err)
		}
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Email lub hasło jest nieprawidłowe",
		})
		return
	}

	if user.Password != password {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Email lub hasło jest nieprawidłowe",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	position, _ := models.GetUserPositionByID(db, user.ID)

	if position == "Magazynowy" {
		c.Redirect(http.StatusFound, "/warehouse/dashboard")
	} else if position == "HR" {
		c.Redirect(http.StatusFound, "/hr/dashboard")
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Twoje stanowisko jest nieprawidłowe",
		})
		return
	}
}
