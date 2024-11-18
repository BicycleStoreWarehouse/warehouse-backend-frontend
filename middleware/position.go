package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"warehouse/models"
)

func WarehouseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Pobranie user_id z sesji
		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Pobranie stanowiska użytkownika z bazy
		userPosition, err := models.GetUserPositionByID(db, userID.(uint))
		if err != nil || userPosition == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Sprawdzenie stanowiska
		if userPosition != "Magazynowy" {
			c.Redirect(http.StatusFound, "/hr/dashboard")
			c.Abort()
			return
		}

		c.Set("user_position", userPosition)
		c.Next()
	}
}

func HrMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Pobranie user_id z sesji
		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Pobranie stanowiska użytkownika z bazy
		userPosition, err := models.GetUserPositionByID(db, userID.(uint))
		if err != nil || userPosition == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Sprawdzenie stanowiska
		if userPosition != "HR" {
			c.Redirect(http.StatusFound, "/warehouse/dashboard")
			c.Abort()
			return
		}

		c.Set("user_position", userPosition)
		c.Next()
	}
}
