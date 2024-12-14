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

		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		userPosition, err := models.GetUserPositionByID(db, userID.(uint))

		if err != nil || userPosition == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		if userPosition != "Magazynowy" && userPosition != "Admin" {
			c.Redirect(http.StatusFound, "/hr/dashboard")
		}

		c.Set("user_position", userPosition)

		c.Next()
	}
}

func HrMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		userPosition, err := models.GetUserPositionByID(db, userID.(uint))

		if err != nil || userPosition == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		if userPosition != "HR" && userPosition != "Admin" {
			c.Redirect(http.StatusFound, "/warehouse/dashboard")
		}

		c.Set("user_position", userPosition)

		c.Next()
	}
}

func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userID := session.Get("user_id")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		userPosition, err := models.GetUserPositionByID(db, userID.(uint))

		if err != nil || userPosition == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		if userPosition != "HR" && userPosition != "Magazynowy" {
			c.Redirect(http.StatusFound, "/admin/dashboard")
		}

		c.Set("user_position", userPosition)

		c.Next()
	}
}
