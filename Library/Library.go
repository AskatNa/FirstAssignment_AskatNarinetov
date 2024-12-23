package main

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[string]Book
}

func (library Library) AddBook(book Book) {
	if _, exists := library.books[book.ID]; exists {
		fmt.Println("Book with such ID already exists.")
		return
	}
	library.books[book.ID] = book
	fmt.Println("Book was added.")
}

func (library Library) BorrowBook(id string) {
	if book, exists := library.books[id]; exists {
		if book.IsBorrowed {
			fmt.Println("Book is borrowed.")
			return
		}
		book.IsBorrowed = true
		library.books[id] = book
		fmt.Println("You have borrowed the book.")
	} else {
		fmt.Println("Book not found.")
	}
}

func (library Library) ReturnBook(id string) {
	if book, exists := library.books[id]; exists {
		if !book.IsBorrowed {
			fmt.Printf("Book is not currently borrowed.")
			return
		}
		book.IsBorrowed = false
		library.books[id] = book
		fmt.Println("You have returned the book")
	} else {
		fmt.Println("Book not found.")
	}
}

func (library Library) ListBooks() {
	for _, book := range library.books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := Library{make(map[string]Book)}
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add book")
		fmt.Println("2.Borrow a book")
		fmt.Println("3. Return a book")
		fmt.Println("4. List of books")
		fmt.Println("5. Exit")
		fmt.Println("Choose and option: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Print("Enter book id: ")
			fmt.Scan(&id)
			fmt.Print("Enter title: ")
			fmt.Scan(&title)
			fmt.Println("Enter an author: ")
			fmt.Scan(&author)

			library.AddBook(Book{ID: id, Title: title, Author: author})

		case 2:
			var id string
			fmt.Println("Enter book id to borrow")
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			var id string
			fmt.Println("Enter book id to return")
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			library.ListBooks()
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
