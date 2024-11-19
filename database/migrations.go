package database

import (
	"time"
	"warehouse/models"

	"gorm.io/gorm"
)

func LoadExampleData(db *gorm.DB) {
	twoYearsAgo := time.Now().AddDate(-2, 0, 0)
	threeYearsAgo := time.Now().AddDate(-3, 0, 0)

	models.CreateUser(db, "Jan", "Kowalski", "jan.kowalski@example.com", "Magazynowy", twoYearsAgo, "123456789", "hashed_password1")
	models.CreateUser(db, "Anna", "Nowak", "anna.nowak@example.com", "Magazynowy", twoYearsAgo, "987654321", "hashed_password2")
	models.CreateUser(db, "Piotr", "Wiśniewski", "piotr.wisniewski@example.com", "Magazynowy", twoYearsAgo, "555555555", "hashed_password3")

	models.CreateUser(db, "Kasia", "Wójcik", "kasia.wojcik@example.com", "HR", threeYearsAgo, "444444444", "hashed_password4")
	models.CreateUser(db, "Tomasz", "Kamiński", "tomasz.kaminski@example.com", "HR", threeYearsAgo, "333333333", "hashed_password5")
	models.CreateUser(db, "Ewa", "Zielińska", "ewa.zielinska@example.com", "Magazynowy", threeYearsAgo, "222222222", "hashed_password6")
}