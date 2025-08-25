package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vidhuuuu/tbr/db"
)

func PrettyPrintBooks(books []db.Book) {
	headings := []string{"ID", "Title", "Author"}
	widths := []int{len(headings[0]), len(headings[1]), len(headings[2])}

	for _, b := range books {
		if l := len(strconv.Itoa(b.ID)); l > widths[0] {
			widths[0] = l
		}
		if l := len(b.Title); l > widths[1] {
			widths[1] = l
		}
		if l := len(b.Author); l > widths[2] {
			widths[2] = l
		}
	}

	printRow := func(cols []string) {
		for i, col := range cols {
			fmt.Printf("%-*s ", widths[i], col)
		}
		fmt.Println()
	}

	printRow(headings)

	for _, w := range widths {
		fmt.Printf("%*s ", w, strings.Repeat("-", w))
	}
	fmt.Println()

	for _, b := range books {
		printRow([]string{
			strconv.Itoa(b.ID),
			b.Title,
			b.Author,
		})
	}
}
