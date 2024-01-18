package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

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

func GiveMoney(db *sql.DB, userID int, sum int) error {
	if _, err := db.Exec(`UPDATE users SET balance = balance + $1 WHERE id = $2`, sum, userID); err != nil {
		return err
	}
	fmt.Printf("Sucessfully money added to %d\n", userID)
	return nil
}