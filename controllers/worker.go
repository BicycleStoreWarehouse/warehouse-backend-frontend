package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

	// Pobierz liczbę urlopów dla użytkownika
	vacationCount, err := models.GetVacationCountByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard_warehouse.html", gin.H{
			"message": "Coś poszło nie tak podczas sprawdzania urlopów",
		})
		return
	}

	// Przygotowanie odpowiedniego komunikatu do szablonu
	var vacationMessage string
	if vacationCount == 1 {
		vacationMessage = "Masz 1 zmianę w zakładce urlop!"
	} else if vacationCount > 1 && vacationCount <= 4 {
		vacationMessage = fmt.Sprintf("Masz %d zmiany w zakładce urlop!", vacationCount)
	} else {
		vacationMessage = fmt.Sprintf("Masz %d zmian w zakładce urlop!", vacationCount)
	}

	// Pobierz liczbę niewykonanych zadań dla użytkownika
	taskCount, err := models.GetUncompletedTasksCountByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard_warehouse.html", gin.H{
			"message": "Coś poszło nie tak podczas sprawdzania zadań",
		})
		return
	}

	// Pobierz liczbę przeterminowanych zadań
	overdueCount, err := models.GetOverdueTasksCountByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard_warehouse.html", gin.H{
			"message": "Coś poszło nie tak podczas sprawdzania przeterminowanych zadań",
		})
		return
	}

	// Przygotowanie odpowiedniego komunikatu do szablonu
	var taskMessage string
	if taskCount == 1 {
		taskMessage = "Masz 1 zadanie do wykonania!"
	} else if taskCount > 1 && taskCount <= 4 {
		taskMessage = fmt.Sprintf("Masz %d zadania do wykonania!", taskCount)
	} else {
		taskMessage = fmt.Sprintf("Masz %d zadań do wykonania!", taskCount)
	}

	// Przygotowanie odpowiedniego komunikatu do szablonu
	var overdueMessage string
	if overdueCount == 1 {
		overdueMessage = "Masz 1 zadanie przeterminowane do wykonania!"
	} else if overdueCount > 1 && overdueCount <= 4 {
		overdueMessage = fmt.Sprintf("Masz %d zadania przeterminowane do wykonania!", overdueCount)
	} else {
		overdueMessage = fmt.Sprintf("Masz %d zadań przeterminowanych do wykonania!", overdueCount)
	}

	c.HTML(http.StatusOK, "dashboard_warehouse.html", gin.H{
		"user_name":          user.Name,
		"task_count":         taskCount,
		"task_message":       taskMessage,
		"overdue_count":      overdueCount,
		"overdue_message":    overdueMessage,
		"has_new_tasks":      taskCount > 0,
		"has_overdue_tasks":  overdueCount > 0,
		"vacation_count":     vacationCount,
		"vacation_message":   vacationMessage,
		"has_vacation_tasks": vacationCount > 0,
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

	// Pobierz dane użytkownika
	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "dashboard_worker.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	// Pobierz urlopy użytkownika
	vacations, err := models.GetVacationsByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard_worker.html", gin.H{
			"message": "Nie udało się pobrać danych o urlopach",
		})
		return
	}

	// Przygotowanie danych o urlopach
	vacationData := []gin.H{}
	for _, vacation := range vacations {
		vacationData = append(vacationData, gin.H{
			"id":         vacation.ID,
			"date_from":  vacation.DateFrom,
			"date_to":    vacation.DateTo,
			"date_count": vacation.DateCount,
			"status":     vacation.Status,
			"reason":     vacation.RejectionReason,
			"read":       vacation.Read,
		})
	}

	// Renderowanie szablonu z listą urlopów
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
		"user_position":    user.Position.Name,
		"user_location":    user.City,
		"vacations":        vacationData,
	})
}

func ChangeStatus(c *gin.Context, db *gorm.DB) {
    // Wydrukuj otrzymane dane dla debugowania
    body, _ := io.ReadAll(c.Request.Body)
    fmt.Printf("Otrzymane dane: %s\n", string(body))
    
    // Przywróć body do kontekstu
    c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

    // Struktura dla danych wejściowych
    var request struct {
        VacationID int `json:"vacation_id" binding:"required"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        fmt.Printf("Błąd parsowania JSON: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error": "Niepoprawne dane wejściowe: " + err.Error(),
        })
        return
    }

    // Aktualizacja statusu
    err := models.UpdateVacationStatus(db, uint(request.VacationID), "Odczytane", nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error": "Nie udało się zaktualizować statusu urlopu",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Status został zaktualizowany",
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
	vacation, err := models.CreateVacation(db, userID.(uint), vacationData.DateFrom, vacationData.DateTo, vacationData.DateCount, "Wysłany")
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
	pdf.Cell(100, 10, fmt.Sprintf("Stanowisko: %s", user.Position.Name))
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

func GetTasks(c *gin.Context, db *gorm.DB) {
	// Sprawdzenie czy użytkownik jest zalogowany
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	// Pobranie zadań dla użytkownika z bazy danych *TYLKO NIEWYKONANYCH*
	tasks, err := models.GetTasksByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "worker_task.html", gin.H{
			"error": "Nie udało się pobrać zadań",
		})
		return
	}

	// Obliczenie liczby dni do daty realizacji dla każdego zadania
	now := time.Now()
	type TaskWithDays struct {
		models.Task
		DaysUntilDeadline int
	}

	var tasksWithDays []TaskWithDays
	for _, task := range tasks {
		daysUntil := int(task.Deadline.Sub(now).Hours() / 24)
		tasksWithDays = append(tasksWithDays, TaskWithDays{
			Task:              task,
			DaysUntilDeadline: daysUntil,
		})
	}

	// Renderowanie szablonu HTML z zadaniami
	c.HTML(http.StatusOK, "worker_task.html", gin.H{
		"Tasks":   tasksWithDays,
		"Message": "Witaj, oto Twoje zadania.",
	})
}

func CompleteTask(c *gin.Context, db *gorm.DB) {
	// Pobierz taskID z formularza
	taskIDStr := c.PostForm("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny identyfikator zadania"})
		return
	}

	// Sprawdzenie czy użytkownik jest zalogowany
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	// Oznaczenie zadania jako wykonanego
	err = models.MarkTaskAsCompleted(db, uint(taskID), userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "worker_task.html", gin.H{
			"error": "Nie udało się oznaczyć zadania",
		})
		return
	}

	// Pobierz zaktualizowaną listę zadań użytkownika
	tasks, err := models.GetTasksByUserID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "worker_task.html", gin.H{
			"error": "Nie udało się pobrać zadań",
		})
		return
	}

	// Oblicz DaysUntilDeadline dla każdego zadania
	for i := range tasks {
		tasks[i].DaysUntilDeadline = int(tasks[i].Deadline.Sub(time.Now()).Hours() / 24)
	}

	// Przekazanie zadań do widoku
	c.HTML(http.StatusOK, "worker_task.html", gin.H{
		"Tasks": tasks,
	})
}

// func ReadVacation(c *gin.Context, db *gorm.DB) {
// 	// Pobierz vacationID z formularza
// 	vacationIDStr := c.PostForm("vacation_id")
// 	vacationID, err := strconv.Atoi(vacationIDStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawny identyfikator urlopu"})
// 		return
// 	}

// 	// Sprawdzenie czy użytkownik jest zalogowany
// 	userID, exists := c.Get("user_id")
// 	if !exists {
// 		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
// 			"message": "Musisz być zalogowany",
// 		})
// 		return
// 	}

// 	// Oznaczenie urlopu jako odczytanego
// 	err = models.MarkVacationAsRead(db, uint(vacationID))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Nie udało się oznaczyć urlopu jako odczytanego",
// 		})
// 		return
// 	}

// 	// Pobierz zaktualizowaną listę urlopów użytkownika
// 	vacations, err := models.GetVacationsByUserID(db, userID.(uint))
// 	if err != nil {
// 		c.HTML(http.StatusInternalServerError, "dashboard_worker.html", gin.H{
// 			"error": "Nie udało się pobrać danych o urlopach",
// 		})
// 		return
// 	}

// 	// Przekazanie urlopów do widoku
// 	c.HTML(http.StatusOK, "dashboard_worker.html", gin.H{
// 		"vacations": vacations,
// 	})
// }
