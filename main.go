package main

import (
	"flag"
	"fmt"
	"time"
)

type TbrItem struct {
	ID     int
	Title  string
	Author string
	Date   time.Time
}

func main() {
	title := flag.String("title", "empty-title", "Title of the book")
	author := flag.String("author", "anonymous", "Author of the book")
	flag.Parse()

	if flag.Parsed() {
		fmt.Printf("title = %s\nauthor = %s\n", *title, *author)
	}
}
