package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HrDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_hr.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	users, _ := models.GetAllUsers(db)

	sales_invoices, purchase_invoices, _ := models.GetPendingInvoices(db)

	var purchase_invoices_data []map[string]interface{}
	for _, invoice := range purchase_invoices {
		purchase_invoice_data := map[string]interface{}{
			"ID":           invoice.ID,
			"SupplierName": invoice.Supplier.Name,
			"NetPrice":     invoice.NetPrice,
			"GrossPrice":   invoice.GrossPrice,
			"StatusName":   invoice.Status.Name,
		}
		purchase_invoices_data = append(purchase_invoices_data, purchase_invoice_data)
	}

	var sales_invoices_data []map[string]interface{}
	for _, invoice := range sales_invoices {
		sales_invoice_data := map[string]interface{}{
			"ID":           invoice.ID,
			"SupplierName": invoice.Customer.Name,
			"NetPrice":     invoice.NetPrice,
			"GrossPrice":   invoice.GrossPrice,
			"StatusName":   invoice.Status.Name,
		}
		sales_invoices_data = append(sales_invoices_data, sales_invoice_data)
	}

	c.HTML(http.StatusOK, "dashboard_hr.html", gin.H{
		"user_name":                 user.Name,
		"users":                     users,
		"pending_purchase_invoices": purchase_invoices_data,
		"pending_sales_invoices":    sales_invoices_data,
	})
}

func ListWorkers(c *gin.Context, db *gorm.DB) {
	users, _ := models.GetAllWorkers(db)

	c.HTML(http.StatusOK, "list_workers.html", gin.H{"users": users})
}

func CreateWorker(c *gin.Context, db *gorm.DB) {
	name := c.DefaultPostForm("name", "")
	surname := c.DefaultPostForm("surname", "")
	email := c.DefaultPostForm("email", "")
	dateOfEmploymentStr := c.DefaultPostForm("date_of_employment", "")
	phone := c.DefaultPostForm("phone", "")
	password := c.DefaultPostForm("password", "")
	street := c.DefaultPostForm("street", "")
	city := c.DefaultPostForm("city", "")
	state := c.DefaultPostForm("state", "")
	zip := c.DefaultPostForm("zip", "")
	country := c.DefaultPostForm("country", "")
	bankAccount := c.DefaultPostForm("bank_account", "")
	nameBank := c.DefaultPostForm("name_bank", "")

	dateOfEmployment, err := time.Parse("2006-01-02", dateOfEmploymentStr)
	if err != nil {
		c.Redirect(http.StatusFound, "/hr/dashboard")
		c.Abort()
		return
	}

	position, _ := models.GetPositionByName(db, "Magazynowy")
	user, err := models.CreateUser(db, name, surname, email, position, dateOfEmployment, phone, password, street, city, state, zip, country, bankAccount, nameBank)

	if err != nil {
		log.Printf("Err %s", err)
		log.Printf("User %s", user.Name)
	}

	ListWorkers(c, db)
}

func DetailWorker(c *gin.Context, db *gorm.DB) {
	userIDParam := c.Param("id")

	userID, _ := strconv.ParseUint(userIDParam, 10, 64)
	userIDUint := uint(userID)

	user, _ := models.GetUserByID(db, userIDUint)

	c.HTML(http.StatusOK, "detail_worker.html", gin.H{"user": user})
}

func UpdateWorkerForm(c *gin.Context, db *gorm.DB) {
	userIDParam := c.Param("id")

	userID, _ := strconv.ParseUint(userIDParam, 10, 64)
	userIDUint := uint(userID)

	user, _ := models.GetUserByID(db, userIDUint)

	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"User": user,
	})
}

func UpdateWorker(c *gin.Context, db *gorm.DB) {
	userIDParam := c.Param("id")

	userID, _ := strconv.ParseUint(userIDParam, 10, 64)
	userIDUint := uint(userID)

	user, _ := models.GetUserByID(db, userIDUint)

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd przy zapisywaniu zmian"})
		return
	}

	ListWorkers(c, db)
}

func ListOrdersHR(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_orders_hr.html", gin.H{"message": "List Orders"})
}

func CreateOrder(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "create_order.html", gin.H{"message": "Create Order"})
}

func ListAplicationsHR(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_applications_hr.html", gin.H{"message": "List Applications"})
}

func DetailApplication(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "detail_application.html", gin.H{"message": "Detail Application"})
}

func UpdateApplication(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "update_application.html", gin.H{"message": "Update Application"})
}

func ShowReport(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "show_report.html", gin.H{"message": "Show Report"})
}
