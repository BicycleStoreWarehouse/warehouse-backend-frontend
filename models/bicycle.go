package models

import "gorm.io/gorm"

type CategoryBicycle struct {
	gorm.Model
	Name     string    `gorm:"unique;not null"`
}

type Bicycle struct {
	gorm.Model
	Name              string `gorm:"not null"`
	CategoryID        uint   `gorm:"not null"`
	Category          CategoryBicycle
	ProducerID        uint 	 `gorm:"not null"`
	Producer          Producer
	Price             float32 `gorm:"not null"`
	Quantity          int     `gorm:"not null"`
	LastQuantityAdded int     `gorm:"not null"`
}

func CreateCategoryBicycle(db *gorm.DB, name string) (CategoryBicycle, error) {
	category := CategoryBicycle{
		Name: name,
	}

	result := db.Create(&category)
	return category, result.Error
}

func CreateBicycle(db *gorm.DB, name string, categoryID, producerID uint, price float32, quantity, lastQuantityAdded int) (Bicycle, error) {
	bicycle := Bicycle{
		Name:              name,
		CategoryID:        categoryID,
		ProducerID:        producerID,
		Price:             price,
		Quantity:          quantity,
		LastQuantityAdded: lastQuantityAdded,
	}

	result := db.Create(&bicycle)
	return bicycle, result.Error
}
