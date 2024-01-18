package main

import (
	"database/sql"
	// . "GoBootcampN8/Month-2/Postgresql/Lesson-8-Indexing/ShopingCenter/"
	_ "github.com/lib/pq"
)


func main() {
	connStr := "user=postgres password=ebot dbname=payuz sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Data
	// users := []User{
	// 	{
	// 		Fname: "Shoping",
	// 		Lname: "Account",
	// 	},
	// 	{
	// 		Fname: "Oybek",
	// 		Lname: "Atamatov",
	// 	},
	// 	{
	// 		Fname: "Doston",
	// 		Lname: "Shernazarov",
	// 	},
	// 	{
	// 		Fname: "Davron",
	// 		Lname: "Nuriddinov",
	// 	},
	// }

	// products := []Product {
	// 	{
	// 		Name: "Bread",
	// 		Price: 3500,
	// 	},
	// 	{
	// 		Name: "Solt",
	// 		Price: 5000,
	// 	},
	// 	{
	// 		Name: "Egg",
	// 		Price: 1200,
	// 	},
	// }

	// Create [-----]
	// CreateUser(db, users[0])
	// CreateUser(db, users[1])
	// CreateUser(db, users[2])
	// CreateUser(db, users[3])
	// CreateProduct(db, products[0])
	// CreateProduct(db, products[1])
	// CreateProduct(db, products[2])

	// Get [-----]
	// GetAllUserInfo(db)
	// GetAllProductInfo(db)

	// Give money to User
	// GiveMoney(db, 1, 5000)
	// GiveMoney(db, 2, 10000)
	// GiveMoney(db, 3, 5000)

	// db userid, productid, companyid
	// UserBuyProduct(db, 2, 1, 1)
	// UserBuyProduct(db, 3, 1, 1)
	// UserBuyProduct(db, 4, 1, 1)
	// UserBuyProduct(db, 4, 2, 1)
	// Get users Money
	
	// GetMoneyUserInfo(db, 1)
	// GetMoneyUserInfo(db, 2)
	// GetMoneyUserInfo(db, 3)
	// GetMoneyUserInfo(db, 4)



	

}