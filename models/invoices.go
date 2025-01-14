package models

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type PurchaseInvoice struct {
	gorm.Model
	SupplierID uint      `gorm:"not null"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID"`
	IssueDate  time.Time `gorm:"not null"`
	NetPrice   float32   `gorm:"not null"`
	GrossPrice float32   `gorm:"not null"`
	StatusID   uint      `gorm:"foreignKey:StatusID"`
	Status     Status    `gorm:"not null"`
	DeliveryID uint      `gorm:"not null"`
	Delivery   Delivery  `gorm:"foreignKey:DeliveryID"`
}

type SalesInvoice struct {
	gorm.Model
	CustomerID uint      `gorm:"not null"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	OrderID    uint      `gorm:"not null"`
	Order      Order     `gorm:"foreignKey:OrderID"`
	IssueDate  time.Time `gorm:"not null"`
	NetPrice   float32   `gorm:"not null"`
	GrossPrice float32   `gorm:"not null"`
	StatusID   uint      `gorm:"foreignKey:StatusID"`
	Status     Status    `gorm:"not null"`
}

func CreateStatus(db *gorm.DB, name string) (Status, error) {
	status := Status{
		Name: name,
	}

	result := db.Create(&status)

	return status, result.Error
}

func CreatePurchaseInvoice(db *gorm.DB, supplierID uint, issueDate time.Time, netPrice float32, status Status, deliveryID uint) (PurchaseInvoice, error) {
	purchaseInvoice := PurchaseInvoice{
		SupplierID: supplierID,
		IssueDate:  issueDate,
		NetPrice:   netPrice,
		GrossPrice: netPrice * 1.23,
		Status:     status,
		DeliveryID: deliveryID,
	}

	result := db.Create(&purchaseInvoice)
	return purchaseInvoice, result.Error
}

func CreateSalesInvoice(db *gorm.DB, customerID, orderID uint, issueDate time.Time, netPrice float32, status Status) (SalesInvoice, error) {
	salesInvoice := SalesInvoice{
		CustomerID: customerID,
		OrderID:    orderID,
		IssueDate:  issueDate,
		NetPrice:   netPrice,
		GrossPrice: netPrice * 1.23,
		Status:     status,
	}

	result := db.Create(&salesInvoice)
	return salesInvoice, result.Error
}

func GetPendingInvoices(db *gorm.DB) ([]SalesInvoice, []PurchaseInvoice, error) {
	var sales_invoices []SalesInvoice
	var purchase_invoices []PurchaseInvoice

	err_sales := db.Preload("Supplier").Preload("Status").Where("status_id = ?", 1).Find(&purchase_invoices).Error

	err_purchase := db.Preload("Customer").Preload("Status").Where("status_id = ?", 1).Find(&sales_invoices).Error

	if err_sales != nil {
		if err_purchase != nil {
			return nil, nil, err_sales
		}
		return nil, purchase_invoices, err_sales
	}

	return sales_invoices, purchase_invoices, nil

}

func CountUnpaidInvoices(db *gorm.DB) (map[string]int64, error) {
    metrics := make(map[string]int64)
    
    // Zmienna tymczasowa dla purchase invoices
    var unpaidPurchase int64
    if err := db.Model(&PurchaseInvoice{}).Where("status_id = ?", 1).Count(&unpaidPurchase).Error; err != nil {
        return nil, err
    }
    metrics["UnpaidPurchaseInvoices"] = unpaidPurchase

    // Zmienna tymczasowa dla sales invoices
    var unpaidSales int64
    if err := db.Model(&SalesInvoice{}).Where("status_id = ?", 1).Count(&unpaidSales).Error; err != nil {
        return nil, err
    }
    metrics["UnpaidSalesInvoices"] = unpaidSales

    return metrics, nil
}
