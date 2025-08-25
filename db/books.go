package db

import (
	"database/sql"
	// "time"
)

// type book struct {
// 	id      int
// 	author  string
// 	title   string
// 	addedAt time.Time
// }

func AddBook(db *sql.DB, author, title string) error {
	query := `INSERT INTO books(author, title) VALUES (?, ?)`
	_, err := db.Exec(query, author, title)
	return err
}
