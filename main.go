package main

import (
	"Library_Management/controllers"
	"Library_Management/models"
	"Library_Management/services"
	"fmt"
)

func main() {
	library := &services.Library{
		Books:   make(map[int]models.Books),
		Members: make(map[int]models.Member),
	}
	fmt.Println("========== Library Management System ==========")
	controllers.BookController(library)
}
