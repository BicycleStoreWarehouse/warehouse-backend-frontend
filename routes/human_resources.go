package routes

import (
	"warehouse/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HumanResourcesRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/dashboard", func(c *gin.Context) {
		controllers.HrDashboard(c, db)
	})
}