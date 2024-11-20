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
	Name       string  `gorm:"not null"`
	Phone      string  `gorm:"not null"`
	Email      string  `gorm:"not null"`
	AddressID  uint
	Address    Address
}

type Producer struct {
	gorm.Model
	Name       string  `gorm:"not null"`
	Country    string  `gorm:"not null"`
	Phone      string  `gorm:"not null"`
	Email      string  `gorm:"not null"`
	AddressID  uint
	Address    Address
	Bicycles   []Bicycle
}

type Client struct {
	gorm.Model
	Name       string  `gorm:"not null"`
	Phone      string  `gorm:"not null"`
	Email      string  `gorm:"not null"`
	AddressID  uint
	Address    Address `gorm:"foreignKey:AddressID"`
}


func CreateAddress(db *gorm.DB, street, city, state, postalCode, country string) (Address, error) {
    address := Address{
        Street:    street,
        City:      city,
        State:     state,
        PostalCode: postalCode,
        Country:   country,
    }

    result := db.Create(&address)
    return address, result.Error
}

func CreateSupplier(db *gorm.DB, name, phone, email string, addressID uint) (Supplier, error) {
    supplier := Supplier{
        Name:      name,
        Phone:     phone,
        Email:     email,
        AddressID: addressID,
    }

    result := db.Create(&supplier)
    return supplier, result.Error
}

func CreateProducer(db *gorm.DB, name, country, phone, email string, addressID uint) (Producer, error) {
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

func CreateClient(db *gorm.DB, name, phone, email string, addressID uint) (Client, error) {
    client := Client{
        Name:      name,
        Phone:     phone,
        Email:     email,
        AddressID: addressID,
    }

    result := db.Create(&client)
    return client, result.Error
}
