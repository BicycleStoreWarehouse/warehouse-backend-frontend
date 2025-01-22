package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

func AdminDashboard(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	usersCount, err := models.CountUsersByPosition(db, nil, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak podczas zliczania użytkowników.",
		})
		return
	}

	incompleteTasks, err := models.CountIncompleteTasks(db, nil, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia nieukończonych zadań.",
		})
		return
	}

	orders, err := models.CountOrders(db, nil, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia zamówień.",
		})
		return
	}

	unpaidInvoices, err := models.CountUnpaidInvoices(db, nil, nil)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia niezapłaconych faktur.",
		})
		return
	}

	usersOnVacation, err := models.CountUsersOnVacation(db)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia niezapłaconych faktur.",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"user_name":                user.Name,
		"pracownik_count":          usersCount["Magazynowy"],
		"hr_count":                 usersCount["HR"],
		"incomplete_tasks":         incompleteTasks,
		"orders_count":             orders,
		"users_vacation":           usersOnVacation,
		"unpaid_purchase_invoices": unpaidInvoices["UnpaidPurchaseInvoices"],
		"unpaid_sales_invoices":    unpaidInvoices["UnpaidSalesInvoices"],
	})

}

func AdminDashboardPost(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err := models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	date1 := c.DefaultPostForm("date_from", "")
    date2 := c.DefaultPostForm("date_to", "")

	if date1 == "" && date2 == "" {
		AdminDashboard(c, db)
		return
	}

	dateFrom, _ := time.Parse("2006-01-02", date1)
    dateTo, _ := time.Parse("2006-01-02", date2)

	usersCount, err := models.CountUsersByPosition(db, &dateFrom, &dateTo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak podczas zliczania użytkowników.",
		})
		return
	}

	incompleteTasks, err := models.CountIncompleteTasks(db, &dateFrom, &dateTo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia nieukończonych zadań.",
		})
		return
	}

	orders, err := models.CountOrders(db, &dateFrom, &dateTo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia zamówień.",
		})
		return
	}

	unpaidInvoices, err := models.CountUnpaidInvoices(db, &dateFrom, &dateTo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia niezapłaconych faktur.",
		})
		return
	}

	usersOnVacation, err := models.CountUsersOnVacation(db)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
			"message": "Błąd podczas liczenia niezapłaconych faktur.",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"user_name":                user.Name,
		"pracownik_count":          usersCount["Magazynowy"],
		"hr_count":                 usersCount["HR"],
		"incomplete_tasks":         incompleteTasks,
		"orders_count":             orders,
		"users_vacation":           usersOnVacation,
		"unpaid_purchase_invoices": unpaidInvoices["UnpaidPurchaseInvoices"],
		"unpaid_sales_invoices":    unpaidInvoices["UnpaidSalesInvoices"],
	})

	userID, exists = c.Get("user_id")
	if !exists {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Musisz być zalogowany",
		})
		return
	}

	user, err = models.GetUserByID(db, userID.(uint))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_dashboard.html", gin.H{
			"message": "Coś poszło nie tak",
		})
		return
	}

	c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
		"user_name": user.Name,
	})
}


func GenerateReport(c *gin.Context, db *gorm.DB) {
    logger := log.New(os.Stdout, "[REPORT-GENERATOR] ", log.LstdFlags)
    
    from := c.DefaultPostForm("date_from", "")
    to := c.DefaultPostForm("date_to", "")
    logger.Printf("Otrzymane daty: from=%s, to=%s", from, to)

    var dateFrom, dateTo *time.Time

    if from != "" && to != "" {
        parsedFrom, err := time.Parse("2006-01-02", from)
        if err == nil {
            dateFrom = &parsedFrom
        }
        
        parsedTo, err := time.Parse("2006-01-02", to)
        if err == nil {
            dateTo = &parsedTo
        }
    }

    // Pobieranie danych z określonego zakresu dat
    usersCount, err := models.CountUsersByPosition(db, dateFrom, dateTo)
    if err != nil {
        logger.Printf("Błąd podczas pobierania liczby użytkowników: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas pobierania danych"})
        return
    }
    
    incompleteTasks, err := models.CountIncompleteTasks(db, dateFrom, dateTo)
    if err != nil {
        logger.Printf("Błąd podczas pobierania nieukończonych zadań: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas pobierania danych"})
        return
    }
    
    orders, err := models.CountOrders(db, dateFrom, dateTo)
    if err != nil {
        logger.Printf("Błąd podczas pobierania zamówień: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas pobierania danych"})
        return
    }
    
    unpaidInvoices, err := models.CountUnpaidInvoices(db, dateFrom, dateTo)
    if err != nil {
        logger.Printf("Błąd podczas pobierania niezapłaconych faktur: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas pobierania danych"})
        return
    }
    
    usersOnVacation, err := models.CountUsersOnVacation(db)
    if err != nil {
        logger.Printf("Błąd podczas pobierania użytkowników na urlopie: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas pobierania danych"})
        return
    }

    logger.Printf("Dane pobrane pomyślnie")

    // Tworzenie nowego dokumentu PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    
    // Używamy Arial zamiast DejaVu
    pdf.SetFont("Arial", "", 12)
    logger.Printf("PDF zainicjalizowany")

    // Nagłówek raportu
    pdf.SetFont("Arial", "B", 16)  // B oznacza pogrubienie
    pdf.CellFormat(190, 10, "Raport statystyk", "", 1, "C", false, 0, "")
    
    pdf.SetFont("Arial", "", 12)
    
    // Formatowanie okresu z obsługą nil
    okresText := "Pelny okres"
    if dateFrom != nil && dateTo != nil {
        okresText = fmt.Sprintf("Okres: %s - %s", dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02"))
    }
    pdf.CellFormat(190, 10, okresText, "", 1, "C", false, 0, "")
    pdf.Ln(5)

    logger.Printf("Nagłówek raportu dodany")

    // Funkcja pomocnicza do dodawania wiersza z danymi
    addRow := func(label string, value interface{}) {
        pdf.SetFillColor(240, 240, 240)
        pdf.CellFormat(130, 8, label, "1", 0, "L", true, 0, "")
        pdf.CellFormat(60, 8, fmt.Sprintf("%v", value), "1", 1, "R", true, 0, "")
    }

    // Nagłówki tabeli
    pdf.SetFont("Arial", "B", 12)  // Pogrubienie dla nagłówków
    pdf.SetFillColor(200, 200, 200)
    pdf.CellFormat(130, 8, "Kategoria", "1", 0, "L", true, 0, "")
    pdf.CellFormat(60, 8, "Wartosc", "1", 1, "R", true, 0, "")

    // Powrót do normalnej czcionki dla danych
    pdf.SetFont("Arial", "", 12)

    // Dane
    addRow("Pracownicy magazynowi", usersCount["Magazynowy"])
    addRow("Pracownicy HR", usersCount["HR"])
    addRow("Nieukonczone zadania", incompleteTasks)
    addRow("Liczba zamowien", orders)
    addRow("Osoby na urlopie", usersOnVacation)
    addRow("Niezaplacone faktury zakupowe", unpaidInvoices["UnpaidPurchaseInvoices"])
    addRow("Niezaplacone faktury sprzedazowe", unpaidInvoices["UnpaidSalesInvoices"])

    logger.Printf("Dane dodane do PDF")

    // Dodanie stopki z datą wygenerowania
    pdf.Ln(10)
    pdf.SetFont("Arial", "", 10)
    pdf.CellFormat(190, 10, fmt.Sprintf("Raport wygenerowany: %s", time.Now().Format("2006-01-02 15:04:05")), "", 1, "L", false, 0, "")

    // Ustawienie nazwy pliku z obsługą nil
    fileName := "raport_pelny_okres.pdf"
    if dateFrom != nil && dateTo != nil {
        fileName = fmt.Sprintf("raport_%s_%s.pdf", 
            dateFrom.Format("2006-01-02"), 
            dateTo.Format("2006-01-02"))
    }
    
    c.Header("Content-Type", "application/pdf")
    c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

    logger.Printf("Próba zapisania PDF do odpowiedzi HTTP")

    // Zapisanie i wysłanie pliku
    err = pdf.Output(c.Writer)
    if err != nil {
        logger.Printf("Błąd podczas generowania PDF: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd podczas generowania raportu"})
        return
    }

    logger.Printf("Raport wygenerowany pomyślnie")
}