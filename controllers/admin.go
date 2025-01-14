package controllers

import (
	"net/http"
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

    usersCount, err := models.CountUsersByPosition(db)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
            "message": "Coś poszło nie tak podczas zliczania użytkowników.",
        })
        return
    }

    incompleteTasks, err := models.CountIncompleteTasks(db)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
            "message": "Błąd podczas liczenia nieukończonych zadań.",
        })
        return
    }

    orders, err := models.CountOrders(db)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
            "message": "Błąd podczas liczenia zamówień.",
        })
        return
    }

    unpaidInvoices, err := models.CountUnpaidInvoices(db)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "admin_dashboard.html", gin.H{
            "message": "Błąd podczas liczenia niezapłaconych faktur.",
        })
        return
    }

    c.HTML(http.StatusOK, "admin_dashboard.html", gin.H{
        "user_name":                 user.Name,
        "pracownik_count":          usersCount["Pracownik"],
        "hr_count":                 usersCount["HR"],
        "incomplete_tasks":         incompleteTasks,
        "orders_count":            orders,
        "unpaid_purchase_invoices": unpaidInvoices["UnpaidPurchaseInvoices"],
        "unpaid_sales_invoices":    unpaidInvoices["UnpaidSalesInvoices"],
    })
}
