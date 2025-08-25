package db

import "time"

type Book struct {
	ID      int
	Title   string
	Author  string
	AddedAt time.Time
	Mode    int
}
