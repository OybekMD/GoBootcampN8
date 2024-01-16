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
	ID    int
	Name  string
	Pages int
}

type Author_Book struct {
	ID int
	Author Author
	Book   Book
}

func main() {
	reqString := []byte(`
		{
			"author" : {
				"name" : "Otkir hoshimov",
				"birthDay" : 1952
			},
			
			"book" : {
				"name": "Otgan kunlar",
				"pages": 300
			}
		}
	`)

	connStr := "user=postgres password=ebot dbname=library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var author_book Author_Book

	if err := json.Unmarshal(reqString, &author_book); err != nil {
		panic(err)
	}

	var respAuthor Author_Book

	rowAuthor := db.QueryRow(`INSERT INTO authors (name, birth_day) VALUES ($1, $2) returning id`, author_book.Author.Name, author_book.Author.BirthDay)

	if err := rowAuthor.Scan(&respAuthor.Author.ID); err != nil {
		panic(err)
	}

	fmt.Println("Author secessfully insterted with id:", respAuthor.Author.ID)

	var respBook Author_Book

	rowBook := db.QueryRow(`INSERT INTO books (name, pages) VALUES ($1, $2) returning id`, author_book.Book.Name, author_book.Book.Pages)
	if err := rowBook.Scan(&respBook.Book.ID); err != nil {
		panic(err)
	}
	fmt.Println("Book secessfully insterted with id:", respBook.Book.ID)

	var respAuthorBook Author_Book
	rowAuthorWithBook := db.QueryRow(`INSERT INTO authors_books (author_id, books_id) VALUES ($1, $2) returning id`, respAuthor.Author.ID, respBook.Book.ID)
	if err := rowAuthorWithBook.Scan(&respAuthorBook.ID); err != nil {
		panic(err)
	}

	fmt.Println("Author Book secessfully insterted with id:", respAuthorBook.ID)
}
