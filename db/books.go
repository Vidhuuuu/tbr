package db

import (
	"database/sql"
	// "time"
)

type Book struct {
	ID      int
	Author  string
	Title   string
}

func AddBook(db *sql.DB, title, author string) error {
	query := `INSERT INTO books(title, author) VALUES (?, ?)`
	_, err := db.Exec(query, title, author)
	return err
}

func ListBooks(db *sql.DB) ([]Book, error) {
	query := `SELECT id, title, author FROM books ORDER BY id ASC`
	rows, err := db.Query(query)
	if err != nil {
	    return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err = rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
