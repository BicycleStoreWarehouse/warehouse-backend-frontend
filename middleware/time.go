package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Dodajemy bazę danych do kontekstu")
		c.Set("db", db)
		c.Next()
	}
}
