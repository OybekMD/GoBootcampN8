package storage

import (
	"database/sql"
	"fmt"
	"handler_test/models"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=ebot dbname=users sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	respUser := models.User{}
	err = db.QueryRow(`INSERT INTO users (id, name, last_name) VALUES($1, $2, $3) RETURNING id, name, last_name`, user.ID, user.FirstName, user.LastName).Scan(
		&respUser.ID,
		&respUser.FirstName,
		&respUser.LastName)
	if err != nil {
		return nil, err
	}

	return &respUser, nil
}

func GetUser(userID string) (*models.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	respUser := models.User{}
	err = db.QueryRow(`SELECT id, name, last_name FROM users WHERE id = $1`, userID).Scan(
		&respUser.ID,
		&respUser.FirstName,
		&respUser.FirstName)
	if err != nil {
		return nil, err
	}

	return &respUser, nil
}

func GetAll(page, limit int) (users []*models.User, err error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	offset := limit * (page - 1)
	rows, err := db.Query(`SELECT id, id_num, name, last_name FROM users LIMIT $1 OFFSET $2`, limit, offset)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.IdNum, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func UpdateUser(user *models.User) (*models.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET name = $1, last_name = $2 where id = $3", user.FirstName, user.LastName, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(userID string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}
	fmt.Printf("Student and associated references successfully deleted with ID: %v\n", userID)
	return nil
}
