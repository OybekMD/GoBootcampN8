package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Author struct {
	ID       int
	Name     string
	BirthDay int
}

type Book struct {
	ID     int
	Name   string
	Pages  int
	Author Author
}

func main() {
	reqString := []byte(`
		{
			"name": "Otgan kunlar",
			"pages": 300,
			"author" : {
				"name" : "Otkir hoshimov",
				"birthDay" : 1952
			}
		}
	`)

	connStr := "user=postgres password= dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var book Book

	if err := json.Unmarshal(reqString, &book); err != nil {
		panic(err)
	}

	var respBook Book

	rowBook := db.QueryRow(`INSERT INTO books (name, pages) VALUES ($1, $2, $3) returning name`, book.Name, book.Pages)

	if err := rowBook.Scan(&respBook.ID, &respBook.Name, &respBook.Pages); err != nil {
		panic(err)
	}

	fmt.Println("Book secessfully insterted with id:", respBook.ID)

	rowAuthor := db.QueryRow(`INSERT INTO author (name, birth_day) VALUES ($1, $2) returning id`, book.Author.Name, book.Author.BirthDay)
	if err := rowAuthor.Scan(&respBook.Author.ID); err != nil {
		panic(err)
	}

	fmt.Println("Author secessfully insterted with id:", respBook.Author.ID)
}
