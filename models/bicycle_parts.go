package models

import (
	"gorm.io/gorm"
)

type CategoryBicycleParts struct {
	gorm.Model
	Name          string          `gorm:"unique;not null"`
	BicycleParts  []BicyclePart   `gorm:"foreignKey:CategoryID"`
}

type BicyclePart struct {
	gorm.Model
	Name             	 string    `gorm:"not null"`
	Type             	 string    
	CategoryID       	 uint      `gorm:"not null"`
	Category        	 CategoryBicycleParts
	ProducerID        	 uint      `gorm:"not null"`        
	Producer             Producer
	Price          	  	 float32   `gorm:"not null"`
	Quantity         	 int       `gorm:"not null"`
	LastQuantityAdded 	 int 	   `gorm:"not null"`
}


func CreateCategoryBicycleParts(db *gorm.DB, name string) (CategoryBicycleParts, error) {
    category := CategoryBicycleParts{
        Name: name,
    }

    result := db.Create(&category)
    return category, result.Error
}

func CreateBicyclePart(db *gorm.DB, name, partType string, categoryID, producerID uint, price float32, quantity, lastQuantityAdded int) (BicyclePart, error) {
    bicyclePart := BicyclePart{
        Name:              name,
        Type:              partType,
        CategoryID:        categoryID,
        ProducerID:        producerID,
        Price:             price,
        Quantity:          quantity,
        LastQuantityAdded: lastQuantityAdded,
    }

    result := db.Create(&bicyclePart)
    return bicyclePart, result.Error
}