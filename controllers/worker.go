package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

func WorkerDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_warehouse.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "dashboard_warehouse.html", gin.H{
		"user_name": user.Name,
	})
}

func DashboardWorker(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_worker.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	vacations, err := models.GetVacationsByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard_worker.html", gin.H{
			"message": "Nie udało się pobrać danych o wakacjach",
		})
		return
	}

	vacationData := []gin.H{}
	for _, vacation := range vacations {
		vacationData = append(vacationData, gin.H{
			"date_from":  vacation.DateFrom,
			"date_to":    vacation.DateTo,
			"date_count": vacation.DateCount,
		})
	}

	c.HTML(http.StatusOK, "dashboard_worker.html", gin.H{
		"user_name":        user.Name,
		"user_surname":     user.Surname,
		"user_email":       user.Email,
		"user_phone":       user.Phone,
		"user_street":      user.Street,
		"user_city":        user.City,
		"user_state":       user.State,
		"user_zip":         user.Zip,
		"user_country":     user.Country,
		"user_bankAccount": user.BankAccount,
		"user_nameBank":    user.NameBank,
		"user_position":    user.PositionID,
		"user_location":    user.City,
		"vacations":        vacationData,
	})
}

func SaveVacation(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	var vacationData struct {
		DateFrom  string `json:"date_from"`
		DateTo    string `json:"date_to"`
		DateCount int    `json:"date_count"`
	}

	if err := c.ShouldBindJSON(&vacationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane wejściowe"})
		return
	}

	// Zapisz nowy urlop do bazy danych
	vacation, err := models.CreateVacation(db, userID.(uint), vacationData.DateFrom, vacationData.DateTo, vacationData.DateCount)
	if err != nil {
		log.Printf("Błąd zapisywania urlopu: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas zapisywania urlopu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Urlop zapisany pomyślnie", "vacation": vacation})
}

func SaveTime(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	var workAndBreakTime struct {
		WorkedHours int `json:"worked_hours"`
		BreakTime   int `json:"break_time"`
	}

	if err := c.ShouldBindJSON(&workAndBreakTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane wejściowe"})
		return
	}

	workedHours := time.Duration(workAndBreakTime.WorkedHours) * time.Second
	breakTime := time.Duration(workAndBreakTime.BreakTime) * time.Second

	err := models.DailyTimekeeping(userID.(uint), workedHours, breakTime, db)

	if err != nil {
		log.Printf("Błąd zapisywania czasu pracy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas zapisywania czasu pracy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Czas pracy zapisany pomyślnie"})
}

func TimeTracking(c *gin.Context, db *gorm.DB) {
	userID, exists_id := c.Get("user_id")

	if !exists_id {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Musisz być zalogowany",
		})
		return
	}

	daily_report, err := models.GetDailyReportForUser(db, userID.(uint))

	if err != nil {
		c.HTML(http.StatusUnauthorized, "time_tracking.html", gin.H{
			"message": "Błąd połączenia z bazą danych",
		})
		return
	}

	c.HTML(http.StatusOK, "time_tracking.html", gin.H{
		"WorkHistory": daily_report,
	})
}

func ListProducts(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_products.html", gin.H{"message": "List Products"})
}

func CreateProduct(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "create_product.html", gin.H{"message": "Create Product"})
}

func ListOrdersWorker(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_orders_worker.html", gin.H{"message": "List Orders"})
}

func DetailOrder(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "detail_order.html", gin.H{"message": "Detail Order"})
}

func UpdateOrder(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "update_order.html", gin.H{"message": "Update Order"})
}

func ListAplicationsWorker(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "list_applications_worker.html", gin.H{"message": "List Applications"})
}

func CreateApplication(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "create_application.html", gin.H{"message": "Create Application"})
}

// GenerateCertificate generates a PDF certificate with the worker's personal information
func GenerateCertificate(c *gin.Context, db *gorm.DB) {
	// Get the user ID from the session
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	// Get the user's data from the database
	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_worker.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 paper size

	// Add a page to the PDF
	pdf.AddPage()

	// Set font for the certificate
	pdf.SetFont("Arial", "B", 16)

	// Title of the certificate
	pdf.Cell(200, 10, "Zaswiadczenie o posiadaniu pracy")

	// Line break
	pdf.Ln(20)

	// Set font for content
	pdf.SetFont("Arial", "", 12)

	// Add personal data to the certificate
	pdf.Cell(190, 10, "Potwierdzamy, ze ponizsza osoba jest zatrudniona w naszej firmie:")
	pdf.Ln(10)
	pdf.Cell(100, 10, fmt.Sprintf("Imie: %s", user.Name))
	pdf.Ln(8)
	pdf.Cell(100, 10, fmt.Sprintf("Nazwisko: %s", user.Surname))
	pdf.Ln(8)
	pdf.Cell(100, 10, fmt.Sprintf("Email: %s", user.Email))
	pdf.Ln(8)
	pdf.Cell(100, 10, fmt.Sprintf("Telefon: %s", user.Phone))
	pdf.Ln(8)
	pdf.Cell(100, 10, fmt.Sprintf("Stanowisko: %s", user.PositionID))
	pdf.Ln(8)
	pdf.Cell(100, 10, fmt.Sprintf("Lokalizacja pracy: %s", user.City))
	pdf.Ln(8)
	pdf.Cell(190, 10, "Data zatrudnienia: "+user.DateOfEmployment.Format("2006-01-02"))
	pdf.Ln(20)
	pdf.Cell(190, 10, "Wygenerowano: "+time.Now().Format("2006-01-02"))
	pdf.Ln(10)

	// Output the PDF to the browser
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=zaswiadczenie.pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Błąd podczas generowania PDF",
		})
	}
}
