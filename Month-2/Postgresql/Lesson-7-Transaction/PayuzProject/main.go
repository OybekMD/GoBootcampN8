package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID      int
	Fname   string
	Lname   string
	Balance int
}

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

	// Create new users
	// CreateUser(db, users[2])

	// Get users Money
	// GetMoneyUser(db, 3)

	// Give money to User
	// GiveMoney(db, 1, 5000)

	// GetAllUserInfo(db)
	// TransferMoney(db, 2, 3, 2000)

}

// GET 
func GetMoneyUser(db *sql.DB, userid int) error {
	var userMoney int
	if err := db.QueryRow(`SELECT u.balance FROM users u WHERE id = $1`, userid).Scan(&userMoney); err != nil {
		return err
	}
	fmt.Printf("Money is %d\n", userMoney)
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

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("Sucessfully money transfered from %d to %d | value: %d\n", userfromID, usertoID, money)
	return nil
}

func CreateUser(db *sql.DB, user User) error {
	var resUser User
	if err := db.QueryRow(`INSERT INTO users(fname, lname) VALUES ($1, $2) 
		RETURNING id, fname, lname`, user.Fname, user.Lname).Scan(&resUser.ID, &resUser.Fname, &resUser.Lname); err != nil {
		return err
	}
	fmt.Printf("User successfully created ! with | %v | %v | %v |\n", resUser.ID, resUser.Fname, resUser.Lname)
	return nil
}
