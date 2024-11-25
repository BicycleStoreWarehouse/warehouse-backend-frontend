package models

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID        uint
    User          User       	 `gorm:"foreignKey:UserID"`
}

type OrderProduct struct {
    gorm.Model
    OrderID        uint
    Order          Order       `gorm:"foreignKey:OrderID"`
    BicycleID      *uint       
    Bicycle        *Bicycle    `gorm:"foreignKey:BicycleID"`
    BicyclePartID  *uint      
	BicyclePart    *BicyclePart `gorm:"foreignKey:BicyclePartID"`
    Quantity       int         
}

func CreateOrder(db *gorm.DB, userID uint) (Order, error) {
    order := Order{
        UserID:        userID,
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
