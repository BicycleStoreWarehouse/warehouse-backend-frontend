package main

import (
	"log"
	"net/http"
	"warehouse/database"
	"warehouse/middleware"
	"warehouse/models"
	"warehouse/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	db.AutoMigrate(&models.User{}, &models.WorkingHoursDaily{}, &models.WorkingHoursMonthly{})
	database.LoadExampleData(db)

	routes.UnauthorizedRoutes(r, db)

	warehouse := r.Group("/warehouse")
	warehouse.Use(middleware.LoginRequiredMiddleware(), middleware.WarehouseMiddleware(db), middleware.DatabaseMiddleware(db))
	{
		routes.WarehouseRoutes(warehouse, db)
		warehouse.POST("/save_time", routes.SaveTime) // Trasa zapisywania czasu pracy
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

		c.Redirect(http.StatusFound, "/login")
	})

	r.POST("/save_time", routes.SaveTime)

	r.Run(":8000")
}

// Funkcja migracji kolumn (zmiana typu danych)
func migrateDB(db *gorm.DB) {
	// Zmiana typu kolumny WorkedHours na VARCHAR(8)
	err := db.Exec("ALTER TABLE working_hours_dailies ALTER COLUMN worked_hours TYPE VARCHAR(8)").Error
	if err != nil {
		log.Fatalf("Nie udało się zmienić typu kolumny WorkedHours: %v", err)
	}

	// Zmiana typu kolumny BreakTime na VARCHAR(8)
	err = db.Exec("ALTER TABLE working_hours_dailies ALTER COLUMN break_time TYPE VARCHAR(8)").Error
	if err != nil {
		log.Fatalf("Nie udało się zmienić typu kolumny BreakTime: %v", err)
	}

	log.Println("Migracje kolumn zakończone pomyślnie.")
}
