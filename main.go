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

	conn, err := db.OpenDB(dsn)
	if err != nil {
	    panic(err)
	}
	defer conn.Close()


	if err = db.InitDB(conn); err != nil {
		panic(err)
	}

	var version string
	err = conn.QueryRow("SELECT sqlite_version()").Scan(&version)
	if err != nil {
	    panic(err)
	}

	fmt.Println("version:", version)
}
