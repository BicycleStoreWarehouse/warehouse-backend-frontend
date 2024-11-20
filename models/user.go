package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string    `gorm:"not null"`
	Surname          string    `gorm:"not null"`
	Email            string    `gorm:"unique;not null"`
	Position         string    `gorm:"not null"`
	DateOfEmployment time.Time `gorm:"not null"`
	Phone            string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
}

func CreateUser(db *gorm.DB, name, surname, email, position string, dateOfEmployment time.Time, phone, password string) (User, error) {

	user := User{
		Name:             name,
		Surname:          surname,
		Email:            email,
		Position:         position,
		DateOfEmployment: dateOfEmployment,
		Phone:            phone,
		Password:         password,
	}

	result := db.Create(&user)

	return user, result.Error
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, gorm.ErrRecordNotFound
		}
		return user, err
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Select("position").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return User{}, gorm.ErrRecordNotFound
		}
		return User{}, err
	}
	return user, nil
}

func GetUserPositionByID(db *gorm.DB, id uint) (position string, error error) {
	var user User

	err := db.Model(&User{}).Select("position").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	return user.Position, nil
}

func GetUsersPassword(db *gorm.DB, email string) (password string, error error) {
	var user User

	err := db.Model(&User{}).Select("password").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	return user.Password, nil
}
