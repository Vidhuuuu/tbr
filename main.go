package main

import (
	"flag"
	"fmt"

	"github.com/Vidhuuuu/tbr/db"
)

func main() {
	title := flag.String("title", "empty-title", "Title of the book")
	author := flag.String("author", "anonymous", "Author of the book")
	flag.Parse()

	fmt.Printf("title = %q, author = %q\n", *title, *author)

	dsn, err := db.PrepareDSN()
	if err != nil {
		panic(err)
	}

	db, err := db.OpenDB(dsn)
	if err != nil {
	    panic(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT sqlite_version()").Scan(&version)
	if err != nil {
	    panic(err)
	}

	fmt.Println("version:", version)
}
