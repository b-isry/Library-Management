package services

import (
	"fmt"
	"../models"
	"../controllers"
)

type Library struct{
	books map[int]models.Books
	members map[int]models.Member
}

type LibraryManager interface{
	AddBook(book Books)
	RemoveBook(bookId int)
	BorrowBook(memberID int, bookID int)
	ReturnBook(memberID int, bookID int)
	ListAvailableBooks() []Books
	ListBorrowedBooks() []Books
	}

func (l *Library) AddBook(book models.Books){
	L.books[book.id] = book
	fmt.Println("Book added successfully")
}

func (l *Library) RemoveBook(bookId int){
	delete(l.books, bookId)
	fmt.Println("Book removed successfully")
}

func (l *Library) BorrowBook(memberID int, bookID int){
	member, ok := l.members[memberID]
	if !ok{
		fmt.Println("Member not found")
		return
	}
	book, ok := l.books[bookID]
	if !ok{
		fmt.Println("Book not found")
		return
	}
	if book.status != "available"{
		fmt.Println("Book is not available")
		return
	}
	book.status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.books[bookID] = book
	fmt.Println("Book borrowed successfully")
}

func (l *Library) ReturnBook(memberID int , bookID int){
	member, ok := l.members[memberID]
	if !ok{
		fmt.Println("Member not found")
		return
	}
	book, ok := l.books[bookID]
	if !ok{
		fmt.Println("Book not found")
		return
	}
	if book.status != "borrowed"{
		fmt.Println("Book is not borrowed")
		return
	}
	book.status = "available"
	member.BorrowedBooks = removeFromSlice(member.BorrowedBooks, book)
	l.books[bookID] = book
	fmt.Println("Book returned successfully")
}

func (l *Library) ListAvailableBooks() []models.Books{
	availableBooks := []models.books
	for _, book := range l.books{
		if book.status == "available"{
			availableBooks = append(availableBooks, book)
		}
	}
	fmt.Println("Available books:")
	for _, book := range availableBooks{
		fmt.Println(book.title)
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks() []models.Books {
	borrowedBooks := []models.Books{}
	for _, book := range l.books {
		if book.status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	fmt.Println("borrowed books:")
	for _, book := range borrowedBooks{
		fmt.Println(book.title)
	}
	return borrowedBooks
}

func removeFromSlice(slice []models.Books, book models.Books) []models.Books{
	for i, b := range slice{
		if b.id == book.id{
			return append(slice[:i], slice[i+1:]...)
		}
	}
}

