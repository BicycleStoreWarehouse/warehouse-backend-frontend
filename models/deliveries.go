package models

import "gorm.io/gorm"

type Delivery struct {
    gorm.Model
    UserID        uint
    User          User       	 `gorm:"foreignKey:UserID"`
}

type DeliveryProduct struct {
    gorm.Model
    OrderID        uint
    Order          Order       `gorm:"foreignKey:OrderID"`
    BicycleID      *uint       
    Bicycle        *Bicycle    `gorm:"foreignKey:BicycleID"`
    BicyclePartID  *uint      
	BicyclePart    *BicyclePart `gorm:"foreignKey:BicyclePartID"`
    Quantity       int         
}


func CreateDelivery(db *gorm.DB, userID uint) (Delivery, error) {
    delivery := Delivery{
        UserID:        userID,
    }

    result := db.Create(&delivery)
    return delivery, result.Error
}

func CreateDeliveryProduct(db *gorm.DB, orderID uint, bicycleID, bicyclePartID *uint, quantity int) (DeliveryProduct, error) {
    deliveryProduct := DeliveryProduct{
        OrderID:        orderID,
        BicycleID:      bicycleID,
        BicyclePartID:  bicyclePartID,
        Quantity:       quantity,
    }

    result := db.Create(&deliveryProduct)
    return deliveryProduct, result.Error
}