package database

import (
	"time"
	"warehouse/models"

	"gorm.io/gorm"
)

func LoadExampleData(db *gorm.DB) {
	twoYearsAgo := time.Now().AddDate(-2, 0, 0)
	threeYearsAgo := time.Now().AddDate(-3, 0, 0)

	position_worker, _ := models.CreatePosition(db, "Magazynowy")
	position_hr, _ := models.CreatePosition(db, "HR")

	worker_1, _ := models.CreateUser(db, "Jan", "Kowalski", "jan.kowalski@example.com", position_worker, twoYearsAgo, "123456789", "hashed_password1", "Chmielna 18", "Warszawa", "Mazowieckie", "05-500", "Polska", "26116055227582461851149641", "Millenium SA")
	worker_2, _ := models.CreateUser(db, "Anna", "Nowak", "anna.nowak@example.com", position_worker, twoYearsAgo, "987654321", "hashed_password2", "Kamienista 92", "Kraków", "Małopolskie", "02-365", "Polska", "30160073920644265811478567", "BNP Paribas Bank SA")
	worker_3, _ := models.CreateUser(db, "Piotr", "Wiśniewski", "piotr.wisniewski@example.com", position_worker, twoYearsAgo, "555555555", "hashed_password3", "Patrice Lumumby 16", "Łódź", "Łódzkie", "91-404", "Polska", "63147013997088337628332906", "Euro Bank SA")
	worker_4, _ := models.CreateUser(db, "Kamil", "Zieliński", "kamil.zielinski@example.com", position_worker, twoYearsAgo, "666666666", "hashed_password4", "Koścista 54", "Poznań", "Wielkopolskie", "83-452", "Polska", "59189088352449104234838308", "Pekao Bank Hipoteczny SA")
	worker_5, _ := models.CreateUser(db, "Marta", "Nowakowska", "marta.nowakowska@example.com", position_worker, twoYearsAgo, "777777777", "hashed_password5", "Sandomierska 20", "Kielce", "Świętokrzyskie", "21-347", "Polska", "64212050770960257003628031", "Santander Consumer Bank SA")

	models.CreateUser(db, "Kasia", "Wójcik", "kasia.wojcik@example.com", position_hr, threeYearsAgo, "444444444", "hashed_password6", "Beskidzka 15", "Bieszczady", "Podkarpackie", "35-540", "Polska", "62188001394953749136465703", "Deutshe Bank Polska SA")
	models.CreateUser(db, "Tomasz", "Kamiński", "tomasz.kaminski@example.com", position_hr, threeYearsAgo, "333333333", "hashed_password7", "Pijna 8", "Katowice", "Śląskie", "25-510", "Polska", "89105040426814845849987809", "ING Bank Śląski SA")
	models.CreateUser(db, "Ewa", "Zielińska", "ewa.zielinska@example.com", position_hr, threeYearsAgo, "222222222", "hashed_password8", "Kolejowa 88", "Wrocław", "Dolnośląskie", "34-460", "Polska", "18114000680270073853609820", "mBank SA")
	models.CreateUser(db, "Karolina", "Wiśniewska", "karolina.wisniewska@example.com", position_hr, threeYearsAgo, "888888888", "hashed_password9", "Centralna 1", "Gdańsk", "Pomorskie", "00-500", "Polska", "78101090585540852529024634", "Narodowy Bank Polski")
	models.CreateUser(db, "Łukasz", "Kowal", "lukasz.kowal@example.com", position_hr, threeYearsAgo, "999999999", "hashed_password10", "Boczna 99", "Piaseczno", "Mazowieckie", "95-530", "Polska", "44249035852272526309101433", "Alior Bank SA")

	address_1, _ := models.CreateAddress(db, "ul. Nowa 1", "Łódź", "Łódzkie", "90-001", "Polska")
	address_2, _ := models.CreateAddress(db, "ul. Kwiatowa 15", "Warszawa", "Mazowieckie", "00-900", "Polska")
	address_3, _ := models.CreateAddress(db, "ul. Słoneczna 12", "Kraków", "Małopolskie", "31-001", "Polska")
	address_4, _ := models.CreateAddress(db, "ul. Zielona 8", "Gdańsk", "Pomorskie", "80-001", "Polska")
	address_5, _ := models.CreateAddress(db, "ul. Szkolna 10", "Poznań", "Wielkopolskie", "60-001", "Polska")

	supplier_1, _ := models.CreateSupplier(db, "Bike Parts Supplier 1", "111223344", "supplier1@example.com", address_1.ID)
	supplier_2, _ := models.CreateSupplier(db, "Bike Parts Supplier 2", "222334455", "supplier2@example.com", address_2.ID)
	supplier_3, _ := models.CreateSupplier(db, "Bike Parts Supplier 3", "333445566", "supplier3@example.com", address_3.ID)
	supplier_4, _ := models.CreateSupplier(db, "Bike Parts Supplier 4", "444556677", "supplier4@example.com", address_4.ID)
	supplier_5, _ := models.CreateSupplier(db, "Bike Parts Supplier 5", "555667788", "supplier5@example.com", address_5.ID)

	producer_1, _ := models.CreateProducer(db, "Shimano", "Japonia", "123456789", "shimano@example.com", address_1.ID)
	producer_2, _ := models.CreateProducer(db, "Specialized", "USA", "987654321", "specialized@example.com", address_2.ID)
	producer_3, _ := models.CreateProducer(db, "Giant", "Tajwan", "456789123", "giant@example.com", address_3.ID)
	producer_4, _ := models.CreateProducer(db, "Trek", "USA", "789123456", "trek@example.com", address_4.ID)
	producer_5, _ := models.CreateProducer(db, "Canyon", "Niemcy", "321654987", "canyon@example.com", address_5.ID)

	customer_1, _ := models.CreateCustomer(db, "Robert", "Lewandowski", "robert.lewandowski@example.com", address_1.ID)
	customer_2, _ := models.CreateCustomer(db, "Magdalena", "Kwiatkowska", "magdalena.kwiatkowska@example.com", address_2.ID)
	customer_3, _ := models.CreateCustomer(db, "Jan", "Nowak", "jan.nowak@example.com", address_3.ID)
	customer_4, _ := models.CreateCustomer(db, "Anna", "Wiśniewska", "anna.wisniewska@example.com", address_4.ID)
	customer_5, _ := models.CreateCustomer(db, "Tomasz", "Zieliński", "tomasz.zielinski@example.com", address_5.ID)

	category_1, _ := models.CreateCategoryBicycleParts(db, "Napęd")
	category_2, _ := models.CreateCategoryBicycleParts(db, "Koła")
	category_3, _ := models.CreateCategoryBicycleParts(db, "Oświetlenie")
	category_4, _ := models.CreateCategoryBicycleParts(db, "Hamulce")
	category_5, _ := models.CreateCategoryBicycleParts(db, "Amortyzatory")

	category_bicycle_1, _ := models.CreateCategoryBicycle(db, "Górski")
	category_bicycle_2, _ := models.CreateCategoryBicycle(db, "Szosa")
	category_bicycle_3, _ := models.CreateCategoryBicycle(db, "Miejski")
	category_bicycle_4, _ := models.CreateCategoryBicycle(db, "Crossowy")
	category_bicycle_5, _ := models.CreateCategoryBicycle(db, "Elektryczny")

	bicyclePart_1, _ := models.CreateBicyclePart(db, "Korba", category_1.ID, producer_1.ID, 150.50, 50, 50)
	bicyclePart_2, _ := models.CreateBicyclePart(db, "Koło 26", category_2.ID, producer_2.ID, 200.00, 30, 30)
	bicyclePart_3, _ := models.CreateBicyclePart(db, "Lampa Przód", category_3.ID, producer_3.ID, 100.00, 20, 15)
	bicyclePart_4, _ := models.CreateBicyclePart(db, "Klocki hamulcowe", category_4.ID, producer_4.ID, 20.00, 100, 80)
	bicyclePart_5, _ := models.CreateBicyclePart(db, "Amortyzator Przód", category_5.ID, producer_5.ID, 500.00, 10, 5)

	bicycle_1, _ := models.CreateBicycle(db, "Mountain Bike", category_bicycle_1.ID, producer_1.ID, 1200.00, 15, 3)
	bicycle_2, _ := models.CreateBicycle(db, "Road Bike", category_bicycle_2.ID, producer_2.ID, 1500.00, 10, 2)
	bicycle_3, _ := models.CreateBicycle(db, "City Cruiser", category_bicycle_3.ID, producer_2.ID, 900.00, 25, 5)
	bicycle_4, _ := models.CreateBicycle(db, "Hybrid Bike", category_bicycle_4.ID, producer_1.ID, 1300.00, 8, 2)
	bicycle_5, _ := models.CreateBicycle(db, "E-Bike", category_bicycle_5.ID, producer_2.ID, 2500.00, 5, 1)

	order_1, _ := models.CreateOrder(db, worker_1.ID)
	order_2, _ := models.CreateOrder(db, worker_2.ID)
	order_3, _ := models.CreateOrder(db, worker_3.ID)
	order_4, _ := models.CreateOrder(db, worker_4.ID)
	order_5, _ := models.CreateOrder(db, worker_5.ID)

	models.CreateOrderProduct(db, order_1.ID, &bicycle_1.ID, nil, 2)
	models.CreateOrderProduct(db, order_1.ID, nil, &bicyclePart_1.ID, 4)
	models.CreateOrderProduct(db, order_2.ID, &bicycle_2.ID, nil, 1)
	models.CreateOrderProduct(db, order_2.ID, nil, &bicyclePart_2.ID, 6)
	models.CreateOrderProduct(db, order_3.ID, &bicycle_3.ID, nil, 3)
	models.CreateOrderProduct(db, order_3.ID, nil, &bicyclePart_3.ID, 2)
	models.CreateOrderProduct(db, order_4.ID, &bicycle_4.ID, nil, 4)
	models.CreateOrderProduct(db, order_4.ID, nil, &bicyclePart_4.ID, 5)
	models.CreateOrderProduct(db, order_5.ID, &bicycle_5.ID, nil, 2)
	models.CreateOrderProduct(db, order_5.ID, nil, &bicyclePart_5.ID, 8)

	delivery_1, _ := models.CreateDelivery(db, worker_1.ID)
	delivery_2, _ := models.CreateDelivery(db, worker_2.ID)
	delivery_3, _ := models.CreateDelivery(db, worker_3.ID)
	delivery_4, _ := models.CreateDelivery(db, worker_4.ID)
	delivery_5, _ := models.CreateDelivery(db, worker_5.ID)

	models.CreateDeliveryProduct(db, delivery_1.ID, nil, &bicyclePart_1.ID, 10)
	models.CreateDeliveryProduct(db, delivery_2.ID, &bicycle_2.ID, nil, 5)
	models.CreateDeliveryProduct(db, delivery_3.ID, nil, &bicyclePart_3.ID, 7)
	models.CreateDeliveryProduct(db, delivery_4.ID, &bicycle_4.ID, nil, 3)
	models.CreateDeliveryProduct(db, delivery_5.ID, nil, &bicyclePart_5.ID, 12)

	status_waiting, _ := models.CreateStatus(db, "Oczekująca")
	status_done, _ := models.CreateStatus(db, "Zrealizowane")

	models.CreatePurchaseInvoice(db, supplier_1.ID, time.Now(), 1000.00, status_waiting, delivery_1.ID)
	models.CreatePurchaseInvoice(db, supplier_2.ID, twoYearsAgo, 2000.00, status_done, delivery_2.ID)
	models.CreatePurchaseInvoice(db, supplier_3.ID, twoYearsAgo, 3000.00, status_done, delivery_2.ID)
	models.CreatePurchaseInvoice(db, supplier_4.ID, time.Now(), 1500.00, status_waiting, delivery_4.ID)
	models.CreatePurchaseInvoice(db, supplier_5.ID, threeYearsAgo, 2500.00, status_done, delivery_5.ID)

	models.CreateSalesInvoice(db, customer_1.ID, order_1.ID, time.Now(), 1200.00, status_done)
	models.CreateSalesInvoice(db, customer_2.ID, order_2.ID, threeYearsAgo, 2200.00, status_done)
	models.CreateSalesInvoice(db, customer_3.ID, order_3.ID, twoYearsAgo, 1800.00, status_done)
	models.CreateSalesInvoice(db, customer_4.ID, order_4.ID, time.Now(), 2500.00, status_waiting)
	models.CreateSalesInvoice(db, customer_5.ID, order_5.ID, threeYearsAgo, 1300.00, status_waiting)

	models.CreateWorkingHoursDaily(db, worker_1.ID, "2024-11-19", "8", "1")
	models.CreateWorkingHoursMonthly(db, worker_1.ID, "2024-11", "160")
}
