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

	r.GET("/workers", func(c *gin.Context) {
		controllers.ListWorkers(c, db)
	})

	r.POST("/workers", func(c *gin.Context) {
		controllers.CreateWorker(c, db)
	})

	r.GET("/create_task", func(c *gin.Context) {
		controllers.CreateTaskHandler(c, db)
	})

	r.POST("/create_task", func(c *gin.Context) {
		controllers.CreateTaskHandler(c, db)
	})

	r.GET("/workers/:id", func(c *gin.Context) {
		controllers.DetailWorker(c, db)
	})

	r.GET("/workers/:id/update", func(c *gin.Context) {
		controllers.UpdateWorkerForm(c, db)
	})

	r.POST("/workers/:id/update", func(c *gin.Context) {
		controllers.UpdateWorker(c, db)
	})

	r.GET("/orders", func(c *gin.Context) {
		controllers.ListOrdersHR(c, db)
	})

	r.POST("/orders", func(c *gin.Context) {
		controllers.CreateOrder(c, db)
	})

	r.GET("/applications", func(c *gin.Context) {
		controllers.ListAplicationsHR(c, db)
	})

	r.GET("/applications/:id", func(c *gin.Context) {
		controllers.DetailApplication(c, db)
	})

	r.PUT("/applications/:id", func(c *gin.Context) {
		controllers.UpdateApplication(c, db)
	})

	r.GET("/report", func(c *gin.Context) {
		controllers.ShowReport(c, db)
	})
}
