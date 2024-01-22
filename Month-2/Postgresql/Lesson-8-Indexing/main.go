package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID      int
	Fname   string
	Lname   string
	Balance int
}

type Product struct {
	ID       int
	Name     string
	Price    int
	CreateAT time.Time
	IsActive bool
}

type Users_Products struct {
	UserID    int
	ProductID int
}

func main() {
	connStr := "user=postgres password=ebot dbname=shopify sslmode=disable"
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
	// GiveMoney(db, 4, 5000)

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

	DeleteUser(db, 2)

	

}

// CREATE
func CreateUser(db *sql.DB, user User) error {
	var resUser User
	if err := db.QueryRow(`INSERT INTO users(fname, lname) VALUES ($1, $2) 
		RETURNING id, fname, lname`, user.Fname, user.Lname).Scan(&resUser.ID, &resUser.Fname, &resUser.Lname); err != nil {
		return err
	}
	fmt.Printf("User successfully created ! with | %v | %v | %v |\n", resUser.ID, resUser.Fname, resUser.Lname)
	return nil
}

func CreateProduct(db *sql.DB, product Product) error {
	var resProduct Product
	if err := db.QueryRow(`INSERT INTO products(name, price) VALUES ($1, $2) 
		RETURNING id, name`,
		product.Name, product.Price).Scan(&resProduct.ID, &resProduct.Name); err != nil {
		return err
	}
	fmt.Printf("Product successfully created ! with | %v | %v | \n", resProduct.ID, resProduct.Name)
	return nil
}

func UserBuyProduct(db *sql.DB, userID int, productID int, companyID int) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return nil
	}

	userMoney, err := GetUserBalance(db, userID) 
	if err != nil {
		tx.Rollback()
		return err
	}

	productPrice, err := GetProductPrice(db, productID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if userMoney < productPrice {
		tx.Rollback()
		fmt.Println("User money less than Product price!")
		return err
	}

	// Subtraction money from User one
	if _, err := tx.Exec(`UPDATE users SET balance = balance - $1 WHERE id = $2`, productPrice, userID); err != nil {
		tx.Rollback()
		return err
	}

	// Add money from User Company
	if _, err := tx.Exec(`UPDATE users SET balance = balance + $1 WHERE id = $2`, productPrice, companyID); err != nil {
		tx.Rollback()
		return err
	}


	if _, err := tx.Exec(`INSERT INTO users_products(user_id, product_id) VALUES ($1, $2)`, userID, productID); err != nil {
		tx.Rollback()
		return err
	}

	// If we don't have any rollback it will commit to our database!
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("Sucessfully product sold! UserID:%d ProductID%d | value: %d\n", userID, productID, productPrice)
	return nil
}

// Read
func GetMoneyUserInfo(db *sql.DB, userid int) error {
	var userMoney int
	if err := db.QueryRow(`SELECT u.balance FROM users u WHERE id = $1`, userid).Scan(&userMoney); err != nil {
		return err
	}
	fmt.Printf("Money is %d\n", userMoney)
	return nil
}
func GetUserBalance(db *sql.DB, userid int) (int, error) {
	var userMoney int
	if err := db.QueryRow(`SELECT u.balance FROM users u WHERE id = $1`, userid).Scan(&userMoney); err != nil {
		return 0, err
	}
	return userMoney, nil
}
func GetProductPrice(db *sql.DB, productid int) (int, error) {
	var pricePrice int
	if err := db.QueryRow(`SELECT p.price FROM products p WHERE id = $1`, productid).Scan(&pricePrice); err != nil {
		return 0, err
	}
	return pricePrice, nil
}

func GetUserInfo(db *sql.DB, id int) error {
	var user User
	err := db.QueryRow("SELECT u.id, u.fname, u.lname, u.balance s.pric birth FROM users u WHERE id = $1", id).Scan(&user.ID, &user.Fname, &user.Lname, &user.Balance)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case err != nil:
		return fmt.Errorf("no User found with ID %d", id)
	default:
		fmt.Println("Info: ", user.ID, user.Fname, user.Lname, user.Balance)
	}
	return nil
}

func GetProductInfo(db *sql.DB, id int) error {
	var product Product
	err := db.QueryRow("SELECT p.id, p.name, p.price, p.created_at, p.isactive birth FROM products p WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.CreateAT, &product.IsActive)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows were returned.")
	case err != nil:
		return fmt.Errorf("no User found with ID %d", id)
	default:
		fmt.Println("Info: ", product.ID, product.Name, product.Price, product.CreateAT, product.IsActive)
	}
	return nil
}

func GetAllUserInfo(db *sql.DB) error {
	rows, err := db.Query("select * from users")
	if err != nil {
		return fmt.Errorf("failed to fetch stundets: %v", err)
	}
	defer rows.Close()
	users := []User{}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Fname, &user.Lname, &user.Balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	for _, u := range users {
		fmt.Println(u.ID, u.Fname, u.Lname, u.Balance)
	}
	return nil
}

func GetAllProductInfo(db *sql.DB) error {
	rows, err := db.Query("select * from products")
	if err != nil {
		return fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreateAT, &product.IsActive)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, product)
	}
	for _, p := range products {
		fmt.Println(p.ID, p.Name, p.Price, p.CreateAT, p.IsActive)
	}
	return nil
}

// UPDATE
func UpdateUser(db *sql.DB, user User) error {
	_, err := db.Exec("UPDATE users SET fname = $1, lname = $2 WHERE id = $3", user.Fname, user.Lname, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update student with ID %d: %v", user.ID, err)
	}
	fmt.Print("User Sucessfuly Updated: ", user)
	return nil
}

func UpdateProduct(db *sql.DB, product Product) error {
	_, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", product.Name, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update student with ID %d: %v", product.ID, err)
	}
	fmt.Print("User Sucessfuly Updated: ", product)
	return nil
}

// DELETE
func DeleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM users_products WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("failed to delete references from student_group for student with ID %d: %v", userID, err)
	}
	_, err = db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return fmt.Errorf("failed to delete students with ID %d: %v", userID, err)
	}

	fmt.Printf("Student and associated references successfully deleted with ID: %d\n", userID)
	return nil
}

func DeleteProduct(db *sql.DB, productID int) error {
	// Deleting product is not correct way to this logic that is why we'll false the isActive!
	_, err := db.Exec("UPDATE products SET isactive = false WHERE id = $1", productID)
	if err != nil {
		return fmt.Errorf("failed to delete students with ID %d: %v", productID, err)
	}

	fmt.Printf("Product successfully deleted with ID: %d\n", productID)
	return nil
}

// OTHER FUNCTION
func GiveMoney(db *sql.DB, userID int, sum int) error {
	if _, err := db.Exec(`UPDATE users SET balance = balance + $1 WHERE id = $2`, sum, userID); err != nil {
		return err
	}
	fmt.Printf("Sucessfully money added to %d\n", userID)
	return nil
}

func TransferMoney(db *sql.DB, userfromID int, usertoID int, money int) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return nil
	}
	// Get money of user one
	var userMoney int
	if err := tx.QueryRow(`SELECT u.balance FROM users u WHERE id = $1`, userfromID).Scan(&userMoney); err != nil {
		tx.Rollback()
		return err
	}
	// If Users money small than Transfer sum it will be break!
	if money > userMoney {
		tx.Rollback()
		return err
	}

	// Subtraction money from User one
	if _, err := tx.Exec(`UPDATE users SET balance = balance - $1 WHERE id = $2`, money, userfromID); err != nil {
		tx.Rollback()
		return err
	}

	// Add money from User one
	if _, err := tx.Exec(`UPDATE users SET balance = balance + $1 WHERE id = $2`, money, usertoID); err != nil {
		tx.Rollback()
		return err
	}
	// If we don't have any rollback it will commit to our database!
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("Sucessfully money transfered from %d to %d | value: %d\n", userfromID, usertoID, money)
	return nil
}
