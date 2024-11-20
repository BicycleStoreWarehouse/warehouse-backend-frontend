package models

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseInvoice struct {
    gorm.Model
    SupplierID    uint       `gorm:"not null"`
    Supplier      Supplier   `gorm:"foreignKey:SupplierID"`
    IssueDate     time.Time  `gorm:"not null"`
    NetPrice      float32    `gorm:"not null"`
    GrossPrice    float32    `gorm:"not null"`
    Status        string     `gorm:"not null"`
    DeliveryID    uint       `gorm:"not null"`
    Delivery      Delivery   `gorm:"foreignKey:DeliveryID"`
}

type SalesInvoice struct {
    gorm.Model
    CustomerID    uint       `gorm:"not null"`
    Customer      Client     `gorm:"foreignKey:CustomerID"`
    OrderID       uint       `gorm:"not null"`
    Order         Order      `gorm:"foreignKey:OrderID"`
    IssueDate     time.Time  `gorm:"not null"`
    NetPrice      float32    `gorm:"not null"`
    GrossPrice    float32    `gorm:"not null"`
    Status        string     `gorm:"not null"`
}

func CreatePurchaseInvoice(db *gorm.DB, supplierID uint, issueDate time.Time, netPrice, grossPrice float32, status string, deliveryID uint) (PurchaseInvoice, error) {
    purchaseInvoice := PurchaseInvoice{
        SupplierID: supplierID,
        IssueDate:  issueDate,
        NetPrice:   netPrice,
        GrossPrice: grossPrice,
        Status:     status,
        DeliveryID: deliveryID,
    }

    result := db.Create(&purchaseInvoice)
    return purchaseInvoice, result.Error
}

func CreateSalesInvoice(db *gorm.DB, customerID, orderID uint, issueDate time.Time, netPrice, grossPrice float32, status string) (SalesInvoice, error) {
    salesInvoice := SalesInvoice{
        CustomerID: customerID,
        OrderID:    orderID,
        IssueDate:  issueDate,
        NetPrice:   netPrice,
        GrossPrice: grossPrice,
        Status:     status,
    }

    result := db.Create(&salesInvoice)
    return salesInvoice, result.Error
}