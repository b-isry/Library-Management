package services

import (
	"Library_Management/models"
	"fmt"
)

type Library struct {
	Books   map[int]models.Books
	Members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Books)
	AddMember(member models.Member)
	RemoveBook(bookId int)
	BorrowBook(memberID int, bookID int)
	ReturnBook(memberID int, bookID int)
	ListAvailableBooks() []models.Books
	ListBorrowedBooks() []models.Books
}

func (l *Library) AddBook(book models.Books) {
	l.Books[book.Id] = book
	book.Status = "available"
	fmt.Println("Book added successfully")
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.Id] = member
	fmt.Println("Member added successfully")
}

func (l *Library) RemoveBook(bookId int) {
	delete(l.Books, bookId)
	fmt.Println("Book removed successfully")
}

func (l *Library) BorrowBook(memberID int, bookID int) {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return
	}
	book, ok := l.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if book.Status != "available" {
		fmt.Println("Book is not available")
		return
	}
	book.Status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Books[bookID] = book

	fmt.Printf(" %s borrowed %s successfully\n", member.Name, book.Title)
}

func (l *Library) ReturnBook(memberID int, bookID int) {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return
	}
	book, ok := l.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if book.Status != "borrowed" {
		fmt.Println("Book is not borrowed")
		return
	}
	book.Status = "available"
	member.BorrowedBooks = removeFromSlice(member.BorrowedBooks, book)
	l.Books[bookID] = book
	fmt.Printf(" %s returned %s successfully\n", member.Name, book.Title)
}

func (l *Library) ListAvailableBooks() []models.Books {
	availableBooks := []models.Books{}
	for _, book := range l.Books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}	
	for _, book := range availableBooks {
		fmt.Println(book.Title)
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks() []models.Books {
	borrowedBooks := []models.Books{}
	for _, book := range l.Books {
		if book.Status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	for _, book := range borrowedBooks {
		fmt.Println(book.Title)
	}
	return borrowedBooks
}

func removeFromSlice(slice []models.Books, book models.Books) []models.Books {
	for i, b := range slice {
		if b.Id == book.Id {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
