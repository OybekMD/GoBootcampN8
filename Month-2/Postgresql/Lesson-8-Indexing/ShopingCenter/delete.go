package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

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