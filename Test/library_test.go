//test the library controller
package Test

import (
	"Library_Management/models"
	"Library_Management/services"
	"testing"
)	

func TestBookController(t *testing.T) {
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
		}

	//test add book
	book := models.Books{
		Id: 1,
		Title: "Test Book",
		Author: "Test Author",
		Status: "available",
		}
	library.AddBook(book)

	//test add member
	member := models.Member{
		Id: 1,
		Name: "Test Member",
		BorrowedBooks: []models.Books{},
		}
	library.AddMember(member)

	//test remove book
	library.RemoveBook(1)

	//test borrow book
	library.BorrowBook(1, 1)

	//test return book
	library.ReturnBook(1, 1)

	//test list available books
	library.ListAvailableBooks()

	//test list borrowed books
	library.ListBorrowedBooks()

}