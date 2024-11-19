package routes

import (
	"warehouse/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WarehouseRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/dashboard", func(c *gin.Context) {
		controllers.WorkerDashboard(c, db)
	})

	r.POST("/save-time", func(c *gin.Context) {
		controllers.SaveTime(c, db)
	})

	r.GET("/time-summary", func(c *gin.Context) {
		controllers.TimeTracking(c, db)
	})
}
