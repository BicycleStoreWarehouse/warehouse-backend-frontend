package main

import (
	"log"
	"net/http"
	"warehouse/database"
	"warehouse/models"
	"warehouse/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	db, err := database.DatabaseConnection()

	db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Błąd połączenia z bazą danych:", err)
	}

	r.LoadHTMLFiles("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusBadRequest, "main_page.go", gin.H{
			"message": "welcome",
		})
	})

	routes.RegisterRoutes(r, db)

	r.Run(":8000")
}