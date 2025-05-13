package main

import (
	"fmt"
	"time"
)

type Book struct {
	title    string
	author   string
	numPages int

	isSaved bool
	savedAt time.Time
}

func (book *Book) saveBook() {
	book.isSaved = true
	book.savedAt = time.Now()
}

func saveBook(book *Book) {
	book.isSaved = true
	book.savedAt = time.Now()

}

func main() {
	var book = Book{
		title:    "",
		author:   "",
		numPages: 0,
		isSaved:  false,
		savedAt:  time.Time{},
	}

	fmt.Println(book)

	book.saveBook()

	fmt.Println(book)
}
