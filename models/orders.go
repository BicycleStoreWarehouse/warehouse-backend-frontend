package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

type OrderProduct struct {
	gorm.Model
	OrderID       uint
	Order         Order `gorm:"foreignKey:OrderID"`
	BicycleID     *uint
	Bicycle       *Bicycle `gorm:"foreignKey:BicycleID"`
	BicyclePartID *uint
	BicyclePart   *BicyclePart `gorm:"foreignKey:BicyclePartID"`
	Quantity      int
}

func CreateOrder(db *gorm.DB, userID uint) (Order, error) {
	order := Order{
		UserID: userID,
	}

	result := db.Create(&order)
	return order, result.Error
}

func CreateOrderProduct(db *gorm.DB, orderID uint, bicycleID, bicyclePartID *uint, quantity int) (OrderProduct, error) {
	orderProduct := OrderProduct{
		OrderID:       orderID,
		BicycleID:     bicycleID,
		BicyclePartID: bicyclePartID,
		Quantity:      quantity,
	}

	result := db.Create(&orderProduct)
	return orderProduct, result.Error
}

func GetAllOrders(db *gorm.DB) ([]OrderProduct, error) {
	var orders []OrderProduct

	if err := db.Preload("Bicycle").Preload("BicyclePart").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func GetUsersOrders(db *gorm.DB, userID uint) ([]OrderProduct, error) {
	var orders []OrderProduct

	err := db.Where("orders.user_id = ?", userID).
		Joins("JOIN orders ON orders.id = order_products.order_id").
		Preload("Bicycle").
		Preload("BicyclePart").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func CountOrders(db *gorm.DB, dateFrom, dateTo *time.Time) (int64, error) {
	var orderCount int64

	// Zbuduj zapytanie z opcjonalnym filtrowaniem dat
	query := db.Model(&Order{})
	if dateFrom != nil && dateTo != nil {
		query = query.Where("created_at BETWEEN ? AND ?", *dateFrom, *dateTo)
	} else if dateFrom != nil {
		query = query.Where("created_at >= ?", *dateFrom)
	} else if dateTo != nil {
		query = query.Where("created_at <= ?", *dateTo)
	}

	// Zlicz zamówienia
	if err := query.Count(&orderCount).Error; err != nil {
		return 0, err
	}

	return orderCount, nil
}


