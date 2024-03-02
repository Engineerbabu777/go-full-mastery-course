package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(title, author string) {
	book := Book{Title: title, Author: author}
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			// Remove the book from the slice.
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Book '%s' by %s removed from the library.\n", book.Title, book.Author)
			return
		}
	}
	fmt.Printf("Book with title '%s' not found in the library.\n", title)
}

func (l *Library) DisplayBooks() {
	fmt.Println("Books in the library:")
	for _, book := range l.Books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}
	fmt.Println()
}

func main() {
	library := Library{}

	library.AddBook("The Great Gatsby", "F. Scott Fitzgerald")
	library.AddBook("To Kill a Mockingbird", "Harper Lee")
	library.AddBook("1984", "George Orwell")

	library.DisplayBooks()

	library.RemoveBook("To Kill a Mockingbird")

	library.DisplayBooks()
}
