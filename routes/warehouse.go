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
}

func DashboardRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/my_dashboard", func(c *gin.Context) {
		controllers.MyDashboard(c, db)
	})
}

func HrRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/my_dashboard_hr", func(c *gin.Context) {
		controllers.MyHrDashboard(c, db)
	})
}
