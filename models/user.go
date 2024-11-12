package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name	  			string		`gorm:"not null"`
	Surname   			string		`gorm:"not null"`
	Email				string		`gorm:"unique;not null"`
	Position 			string		`gorm:"not null"`
	DateOfEmployment	time.Time	`gorm:"not null"`
	Phone				string		`gorm:"unique;not null"`
	Password           	string    	`gorm:"not null"`
}

func CreateUser(db *gorm.DB,
	name, surname, email, position string, dateOfEmployment time.Time, phone, password string) error {

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

    return result.Error
}

func GetUser(db *gorm.DB, email string) (user User, err error) {

	err = db.Model(&User{}).Select("position").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return user, err
	}

	return user, nil
}


func GetUserPosition(db *gorm.DB, email string) (position string, error error) {
	var user User

	err := db.Model(&User{}).Select("position").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	return user.Position, nil
}

func GetUserPassword(db *gorm.DB, email string) (password string, error error) {
	var user User

	err := db.Model(&User{}).Select("password").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "not found", nil
		}
		return "", err
	}

	return user.Password, nil
}

func GetUserNameAndSurname(db *gorm.DB, email string) (name string, surname string, error error) {
	var user User

	err := db.Model(&User{}).Select("name", "surname").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", nil
		}
		return "", "", err
	}

	return user.Name, user.Surname, nil
}