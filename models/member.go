package models

type Member struct{
	id int
	name string
	BorrowedBooks []Books
}