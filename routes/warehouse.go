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

	r.GET("/generate-certificate", func(c *gin.Context) {
		controllers.GenerateCertificate(c, db)
	})

	r.GET("/my_task", func(c *gin.Context) {
		controllers.GetTasks(c, db)
	})

	r.POST("/my_task", func(c *gin.Context) {
		controllers.CompleteTask(c, db)
	})

	r.GET("/dashboard-worker", func(c *gin.Context) {
		controllers.DashboardWorker(c, db)
	})

	r.POST("/save-vacation", func(c *gin.Context) {
		controllers.SaveVacation(c, db)
	})

	r.GET("/time-summary", func(c *gin.Context) {
		controllers.TimeTracking(c, db)
	})

	r.GET("/products", func(c *gin.Context) {
		controllers.ListProducts(c, db)
	})

	r.POST("/products", func(c *gin.Context) {
		controllers.CreateProduct(c, db)
	})

	r.GET("/orders", func(c *gin.Context) {
		controllers.ListOrdersWorker(c, db)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		controllers.DetailOrder(c, db)
	})

	r.PUT("/orders/:id", func(c *gin.Context) {
		controllers.UpdateOrder(c, db)
	})

	r.GET("/applications", func(c *gin.Context) {
		controllers.ListAplicationsWorker(c, db)
	})

	r.POST("/applications", func(c *gin.Context) {
		controllers.CreateApplication(c, db)
	})

}
