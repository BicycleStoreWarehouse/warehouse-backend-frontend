package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street     string `gorm:"not null"`
	City       string `gorm:"not null"`
	State      string
	PostalCode string `gorm:"not null"`
	Country    string `gorm:"not null"`
}

type Supplier struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null"`
	AddressID uint
	Address   Address
}

type Producer struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Country   string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null"`
	AddressID uint
	Address   Address
}

type Customer struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Phone     string `gorm:"not null"`
	Email     string `gorm:"not null"`
	AddressID uint
	Address   Address `gorm:"foreignKey:AddressID"`
}

func CreateAddress(db *gorm.DB, street string, city string, state string, postalCode string, country string) (Address, error) {
	address := Address{
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
		Country:    country,
	}

	result := db.Create(&address)
	return address, result.Error
}

func CreateSupplier(db *gorm.DB, name string, phone string, email string, addressID uint) (Supplier, error) {
	supplier := Supplier{
		Name:      name,
		Phone:     phone,
		Email:     email,
		AddressID: addressID,
	}

	result := db.Create(&supplier)
	return supplier, result.Error
}

func CreateProducer(db *gorm.DB, name string, country string, phone string, email string, addressID uint) (Producer, error) {
	producer := Producer{
		Name:      name,
		Country:   country,
		Phone:     phone,
		Email:     email,
		AddressID: addressID,
	}

	result := db.Create(&producer)
	return producer, result.Error
}

func CreateCustomer(db *gorm.DB, name string, phone string, email string, addressID uint) (Customer, error) {
	client := Customer{
		Name:      name,
		Phone:     phone,
		Email:     email,
		AddressID: addressID,
	}

	result := db.Create(&client)
	return client, result.Error
}
