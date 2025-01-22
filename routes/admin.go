package routes

import (
	"warehouse/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/dashboard", func(c *gin.Context) {
		controllers.AdminDashboard(c, db)
	})
	
	r.POST("/dashboard", func(c *gin.Context) {
		controllers.AdminDashboardPost(c, db)
	})

	r.POST("/generate-report", func(c *gin.Context) {
		controllers.GenerateReport(c, db)
	})
}
