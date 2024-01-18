package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

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