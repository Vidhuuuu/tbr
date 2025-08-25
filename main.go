package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/Vidhuuuu/tbr/db"
	"github.com/Vidhuuuu/tbr/utils"
)

func main() {
	list := flag.Bool("list", false, "list tbr")
	book := flag.String("add", "empty by empty", "add to tbr")
	del := flag.String("del", "empty", "delete from storage")
	flag.Parse()

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

	if *list {
		books, err := db.ListBooks(conn)
		if err != nil {
			panic(err)
		}

		if len(books) == 0 {
			fmt.Println("empty tbr")
			return
		}
		utils.PrettyPrintBooks(books)
		return
	}

	if *del != "empty" {
		parts := strings.SplitSeq(*del, " ")
		for p := range parts {
			if p == "" {
				continue
			}
			id, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			db.DeleteBook(conn, id)
		}
		return
	}

	parts := strings.Split(*book, " by ")
	if len(parts) != 2 {
		panic(fmt.Errorf("inconsistent name: %v\n", *book))
	}

	title := parts[0]
	author := parts[1]

	err = db.AddBook(conn, title, author)
	if err != nil {
		panic(err)
	}
	fmt.Printf("added %s by %s\n", title, author)
}
