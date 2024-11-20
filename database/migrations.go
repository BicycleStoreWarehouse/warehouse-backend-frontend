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
	user, _ := models.CreateUser(db, "Ewa", "Zielińska", "ewa.zielinska@example.com", "Magazynowy", threeYearsAgo, "222222222", "hashed_password6")

	models.CreateAddress(db, "ul. Nowa 1", "Łódź", "Łódzkie", "90-001", "Polska")
	models.CreateAddress(db, "ul. Kwiatowa 15", "Warszawa", "Mazowieckie", "00-900", "Polska")

	models.CreateSupplier(db, "Bike Parts Supplier", "111223344", "supplier@example.com", 1)

	models.CreateProducer(db, "Shimano", "Japonia", "123456789", "shimano@example.com", 1)
	models.CreateProducer(db, "Specialized", "USA", "987654321", "specialized@example.com", 2)

	models.CreateClient(db, "Robert", "Lewandowski", "robert.lewandowski@example.com", 1)
	models.CreateClient(db, "Magdalena", "Kwiatkowska", "magdalena.kwiatkowska@example.com", 2)

	models.CreateCategoryBicycleParts(db, "Napęd")
	models.CreateCategoryBicycleParts(db, "Koła")
	models.CreateCategoryBicycleParts(db, "Oświetlenie")

	models.CreateCategoryBicycle(db, "Górski")
	models.CreateCategoryBicycle(db, "Szosa")

	models.CreateBicyclePart(db, "Korba", "Napęd", 1, 1, 150.50, 50, 10)
	bicyclePart, _ := models.CreateBicyclePart(db, "Koło 26", "Koła", 2, 2, 200.00, 30, 5)

	models.CreateBicycle(db, "Mountain Bike", 1, 1, 1200.00, 15, 3)
	bicycle, _ := models.CreateBicycle(db, "Road Bike", 2, 2, 1500.00, 10, 2)

	order, _ := models.CreateOrder(db, user.ID)

	models.CreateOrderProduct(db, order.ID, &bicycle.ID, &bicyclePart.ID, 2)

	models.CreateDelivery(db, user.ID)

	models.CreateDeliveryProduct(db, order.ID, &bicycle.ID, &bicyclePart.ID, 10)

	models.CreatePurchaseInvoice(db, 1, time.Now(), 1000.00, 1230.00, "Oczekująca", 1)

	models.CreateSalesInvoice(db, 1, 1, time.Now(), 1200.00, 1500.00, "Zrealizowane")

	models.CreateWorkingHoursDaily(db, 1, "2024-11-19", "8", "1")
	models.CreateWorkingHoursMonthly(db, 1, "2024-11", "160")
}
