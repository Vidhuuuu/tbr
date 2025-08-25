package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Vidhuuuu/tbr/db"
	"github.com/Vidhuuuu/tbr/utils"
)

func main() {
	args := os.Args[1:]

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

	switch args[0] {
	case "list":
		books, err := db.ListBooks(conn)
		if err != nil {
			panic(err)
		}

		if len(books) == 0 {
			fmt.Println("Empty TBR")
			return
		}
		utils.PrettyPrintBooks(books)
	case "add":
		raw := strings.Join(args[1:], " ")
		parts := strings.Split(raw, " by ")
		if len(parts) != 2 {
			panic(fmt.Errorf("inconsistent name: %s\n", raw))
		}

		title := parts[0]
		author := parts[1]

		if err = db.AddBook(conn, title, author); err != nil {
			panic(err)
		}
		fmt.Printf("added %s by %s\n", title, author)
	case "del":
		if len(args) < 2 {
			fmt.Println("Usage: tbr del <id> [id2, id3, ...]")
			return
		}
		for _, a := range args[1:] {
			id, err := strconv.Atoi(a)
			if err != nil {
				fmt.Printf("invalid id: %s\n", a)
				continue
			}
			db.DeleteBook(conn, id)
		}
	default:
		fmt.Println("unknown command:", args[0])
		fmt.Println("Usage: tbr [list|add|del] ...")
		return
	}
}
