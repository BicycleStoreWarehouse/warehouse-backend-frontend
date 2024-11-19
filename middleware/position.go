package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"warehouse/models"
)

func PositionMiddleware(db *gorm.DB) gin.HandlerFunc {
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

	}
}
