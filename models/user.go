package models

import (
	"go.starlark.net/lib/time"
    "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name	  			string		`gorm:"not null"`
	Surname   			string		`gorm:"not null"`
	Email				string		`gorm:"unique;not null"`
	Position 			string		`gorm:"not null"`
	DateOfEmployment	time.Time	`gorm:"not null"`
	Phone				string		`gorm:"not null"`
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

func UserExists(db *gorm.DB, email string) (bool, error) {
    var exists bool

    err := db.Model(&User{}).
        Select("1").
        Where("email = ?", email).
        Limit(1).
        Find(&exists).Error

    if err != nil && err != gorm.ErrRecordNotFound {
        return false, err
    }
    return exists, nil
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
			return "", nil
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