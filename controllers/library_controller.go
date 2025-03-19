package controllers

import (
	"fmt"
	"github.com/yourusername/projectname/models"
	"github.com/yourusername/projectname/services"
)

func bookController(library *services.Library) {
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")

	var choice int
	fmt.Scan(&choice)

	switch choice{
	case 1:
		var book models.Books
		fmt.Println("Enter book title: ")
		fmt.Scan(&book.Title)
		fmt.Println("Enter book author: ")
		fmt.Scan(&book.Author)
		addBook(library, book)
	case 2:
		var bookId int
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookId)
		removeBook(library, bookId)
	case 3:
		var memberID int
		var bookID int
		fmt.Println("Enter member id: ")
		fmt.Scan(&memberID)
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookID)
		borrowBook(library, memberID, bookID)
	case 4:
		var memberID int
		var bookID int
		fmt.Println("Enter member id: ")
		fmt.Scan(&memberID)
		fmt.Println("Enter book id: ")
		fmt.Scan(&bookID)
		returnBook(library, memberID, bookID)
	case 5:
		fmt.Println("Available books: ")
		listAvailableBooks(library)
	case 6:
		fmt.Println("Borrowed books: ")
		listBorrowedBooks(library)
	default:
		fmt.Println("Invalid choice")
	}

}




