package main

import (
	"fmt"

	"github.com/owen0163/pack/book"
)

func main() {
	var b book.Book
	b.Title = "A Book"
	b.Author = "Pallat"

	fmt.Println(b)
}
