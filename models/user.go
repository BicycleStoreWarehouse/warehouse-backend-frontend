package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Surname          string `gorm:"not null"`
	Email            string `gorm:"unique;not null"`
	PositionID       uint
	Position         Position  `gorm:"foreignKey:PositionID"`
	DateOfEmployment time.Time `gorm:"not null"`
	Phone            string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
	Street           string    `gorm:"not null"`
	City             string    `gorm:"not null"`
	State            string    `gorm:"not null"`
	Zip              string    `gorm:"not null"`
	Country          string    `gorm:"not null"`
	BankAccount      string    `gorm:"unique;not null"`
	NameBank         string    `gorm:"not null"`
}

func CreateUser(db *gorm.DB, name string, surname string, email string, position Position, dateOfEmployment time.Time, phone, password, street, city, state, zip, country, bankAccount, nameBank string) (User, error) {

	user := User{
		Name:             name,
		Surname:          surname,
		Email:            email,
		Position:         position,
		DateOfEmployment: dateOfEmployment,
		Phone:            phone,
		Password:         password,
		Street:           street,
		City:             city,
		State:            state,
		Zip:              zip,
		Country:          country,
		BankAccount:      bankAccount,
		NameBank:         nameBank,
	}

	result := db.Create(&user)

	return user, result.Error
}

func CreatePosition(db *gorm.DB, name string) (Position, error) {
	position := Position{
		Name: name,
	}

	result := db.Create(&position)

	return position, result.Error
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var user User
	err := db.Preload("Position").Where("id = ?", id).First(&user).Order("id DESC").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return User{}, gorm.ErrRecordNotFound
		}
		return User{}, err
	}

	result := db.Create(&user)

	return user, result.Error
}

func CreatePosition(db *gorm.DB, name string) (Position, error) {
	position := Position{
		Name: name,
	}

	result := db.Create(&position)

	return position, result.Error
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var user User
	err := db.Preload("Position").Where("id = ?", id).First(&user).Order("id DESC").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return User{}, gorm.ErrRecordNotFound
		}
		return User{}, err
	}
	return user, nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User

	err := db.Find(&users).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return users, nil
}

func GetAllWorkers(db *gorm.DB) ([]User, error) {
	var users []User

	err := db.Joins("JOIN positions ON users.position_id = positions.id").
			Where("positions.name = ?", "Magazynowy").
			Find(&users).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return users, nil
}

func GetUserPositionByID(db *gorm.DB, id uint) (name string, error error) {
	var position Position

	err := db.Model(&Position{}).
		Select("positions.name").
		Joins("JOIN users ON users.position_id = positions.id").
		Where("users.id = ?", id).
		First(&position).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	return position.Name, nil
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

func GetPositionByName(db * gorm.DB, name string) (Position, error) {
	var position Position

	err := db.Where("name = ?", name).First(&position).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Position{}, gorm.ErrRecordNotFound
		}
		return Position{}, err
	}

	return position, err
}

func CountUsersByPosition(db *gorm.DB) (map[string]int64, error) {
    var users []User

    if err := db.Preload("Position").Find(&users).Error; err != nil {
        return nil, err
    }

    var workerCount int64
    var hrCount int64

    for _, user := range users {
        switch user.Position.Name {
        case "Magazynowy":
            workerCount++
        case "HR":
            hrCount++
        }
    }

    return map[string]int64{
        "Magazynowy": workerCount,
        "HR": hrCount,
    }, nil
}

