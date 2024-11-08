package main

import (
	"log"
	"net/http"
	"warehouse/controllers"
	"warehouse/database"
	"warehouse/middleware"
	"warehouse/models"
	"warehouse/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	store := cookie.NewStore([]byte("secret"))

	r := gin.Default()
	r.Static("/styles", "./styles")
	r.LoadHTMLGlob("templates/*")
	r.Use(sessions.Sessions("warehouse", store))

	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("Błąd połączenia z bazą danych:", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.User{})

	routes.UnauthorizedRoutes(r, db)

	warehouse := r.Group("/warehouse")
	warehouse.Use(middleware.LoginRequiredMiddleware(), middleware.WarehouseMiddleware(db))
	{
		routes.WarehouseRoutes(warehouse, db)
	}

	hr := r.Group("/hr")
	hr.Use(middleware.LoginRequiredMiddleware(), middleware.HrMiddleware(db))
	{
		routes.HumanResourcesRoutes(hr, db)
	}

	r.GET("/logout", middleware.LoginRequiredMiddleware(), func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Redirect(http.StatusFound, "/home")
	})

	r.Run(":8000")
}
