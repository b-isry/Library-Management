# Library Management System

A simple command-line library management system written in Go that allows librarians to manage books and members.

## Features

- Add new books to the library
- Add new members
- Remove books from the library
- Borrow books
- Return books
- List available books
- List borrowed books

## Getting Started

### Prerequisites

- Go 1.x or higher
- Git (optional)

### Installation

1. Clone the repository (or download the source code):

```bash
git clone https://github.com/yourusername/Library-Management.git
cd Library-Management
```

2. Build the project:

```bash
go build
```

3. Run the application:

```bash
./Library-Management
```

## Usage

The system presents a menu with the following options:

1. **Add Book**
   - Enter book ID (integer)
   - Enter book title
   - Enter book author
   - Book will be marked as "available" automatically

2. **Add Member**
   - Enter member ID (integer)
   - Enter member name
   - Member will be initialized with empty borrowed books list

3. **Remove Book**
   - Enter book ID to remove it from the system

4. **Borrow Book**
   - Enter member ID
   - Enter book ID
   - System will check if:
     - Member exists
     - Book exists
     - Book is available

5. **Return Book**
   - Enter member ID
   - Enter book ID
   - System will check if:
     - Member exists
     - Book exists
     - Book is currently borrowed

6. **List Available Books**
   - Displays all books with "available" status

7. **List Borrowed Books**
   - Displays all books with "borrowed" status

8. **Exit**
   - Exits the program

After each operation, the system will ask if you want to continue (y/n).

## Project Structure

```
Library-Management/
├── main.go                 # Entry point
├── controllers/
│   └── library_controller.go  # User interface handling
├── models/
│   └── models.go          # Data structures
└── services/
    └── library_service.go # Business logic

```

## Data Structures

### Book

- Id: integer
- Title: string
- Author: string
- Status: string ("available" or "borrowed")

### Member

- Id: integer
- Name: string
- BorrowedBooks: []Books

## Error Handling

The system handles various error cases:

- Member not found
- Book not found
- Book not available for borrowing
- Book not borrowed (when trying to return)
- Invalid menu choices

## Contributing

Feel free to submit issues and enhancement requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
