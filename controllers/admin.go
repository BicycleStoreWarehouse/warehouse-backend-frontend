package controllers

import (
	"net/http"
	"time"
	"warehouse/models"

	"github.com/gin-gonic/gin"
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
