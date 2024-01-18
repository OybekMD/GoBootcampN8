package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

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