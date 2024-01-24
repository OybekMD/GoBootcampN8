package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"nt_bootcamp/bootcamp_go/gin/models"
)

func connect() (*sql.DB, error) {
	dsn := "user=doston password=doston dbname=najot_talim sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func CreateUser(user *models.UserGin) (*models.UserGin, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	respUser := models.UserGin{}
	err = db.QueryRow(`INSERT INTO users (id, first_name, last_name) VALUES ($1, $2, $3) returning id, first_name, last_name`,
		user.ID,
		user.FirstName,
		user.LastName).Scan(&respUser.ID, &respUser.FirstName, &respUser.LastName)
	if err != nil {
		return nil, err
	}

	return &respUser, nil

}

func GetUser(userID string) (*models.UserGin, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	respUser := models.UserGin{}
	err = db.QueryRow(`SELECT id, first_name,last_name FROM users WHERE id=$1`, userID).Scan(
		&respUser.ID,
		&respUser.FirstName,
		&respUser.LastName)
	if err != nil {
		return nil, err
	}

	return &respUser, nil
}

func GetAll(page, limit int) (users []*models.UserGin, err error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	offset := limit * (page - 1)
	rows, err := db.Query(`SELECT id, first_name, last_name FROM users LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.UserGin
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func UpdateUserById(user *models.UserGin) (*models.UserGin, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET first_name = $1, last_name = $2 where id = $3", user.FirstName, user.LastName, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserByID(userID string) error {
	db, err := connect()
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
