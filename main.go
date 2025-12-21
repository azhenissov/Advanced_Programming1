package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/azhenissov/Advanced_Programming1/Bank"
)

func main() {

	/*
		-------------------Shapes----------
	*/
	//	//shapes := []Shapes.Shape{
	//	//	Shapes.NewRectangle(4, 5),
	//	//	Shapes.NewCircle(3),
	//	//	Shapes.NewSquare(5),
	//	//	Shapes.NewTriangle(3, 4, 5),
	//	//}
	//	//
	//	//for i, shape := range shapes {
	//	//	fmt.Printf("Shape: %d\n", i+1)
	//	//	fmt.Printf("Area: %.2f\n", shape.Area())
	//	//	fmt.Printf("Perimeter: %.2f\n\n", shape.Perimeter())
	//	//}
	//
	/*
		-------------------- Library ----------------
	*/
	//	lib := &Library.Library{}
	//	reader := bufio.NewReader(os.Stdin)
	//
	//	for {
	//		fmt.Println("===== Library Menu ===== ")
	//		fmt.Println("1.Add Book")
	//		fmt.Println("2.Borrow Book")
	//		fmt.Println("3.Return Book")
	//		fmt.Println("4.List Available Books")
	//		fmt.Println("5.Exit")
	//
	//		input, _ := reader.ReadString('\n')
	//		input = strings.TrimSpace(input)
	//
	//		switch input {
	//		case "1":
	//			addBook(lib, reader)
	//		case "2":
	//			borrowBook(lib, reader)
	//		case "3":
	//			returnBook(lib, reader)
	//		case "4":
	//			listBooks(lib)
	//		case "5":
	//			fmt.Println("Bye Bye")
	//			return
	//		default:
	//			fmt.Println("Invalid Input")
	//		}
	//
	//		fmt.Println()
	//	}
	//
	//}
	//
	//func addBook(lib *Library.Library, reader *bufio.Reader) {
	//	fmt.Println("Add Book ID")
	//	idStr, _ := reader.ReadString('\n')
	//	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	//
	//	fmt.Println("Add Book Title")
	//	title, _ := reader.ReadString('\n')
	//
	//	fmt.Println("Add book Author")
	//	author, _ := reader.ReadString('\n')
	//
	//	book := Library.Book{
	//		ID:         id,
	//		Title:      strings.TrimSpace(title),
	//		Author:     strings.TrimSpace(author),
	//		IsBorrowed: false,
	//	}
	//	lib.AddBook(book)
	//	fmt.Println("Book added successfully")
	//}
	//
	//func borrowBook(lib *Library.Library, reader *bufio.Reader) {
	//	fmt.Println("Borrow Book by ID")
	//	idStr, _ := reader.ReadString('\n')
	//	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	//
	//	book, err := lib.BorrowBook(id)
	//	if err != nil {
	//		fmt.Println("Error: ", err)
	//		return
	//	}
	//
	//	fmt.Printf("You borrowed book %s by %s\n", book.Title, book.Author)
	//}
	//
	//func returnBook(lib *Library.Library, reader *bufio.Reader) {
	//	fmt.Println("Return Book by ID")
	//	idStr, _ := reader.ReadString('\n')
	//	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	//
	//	err := lib.ReturnBook(id)
	//	if err != nil {
	//		fmt.Println("Error: ", err)
	//		return
	//	}
	//	fmt.Println("Book Returned successfully")
	//}
	//func listBooks(lib *Library.Library) {
	//	books, err := lib.ListAvailableBooks()
	//
	//	if err != nil {
	//		fmt.Println("Error: ", err)
	//		return
	//	}
	//
	//	if len(books) == 0 {
	//		fmt.Println("No books found")
	//	}
	//
	//	fmt.Println("Available Books:")
	//	for _, book := range books {
	//		fmt.Printf("ID: %d | %s by %s\n", book.ID, book.Title, book.Author)
	//	}
	//
	/*
		---------------Company-----------------
	*/
	//company := Company.NewCompany()
	//
	//company.AddEmployee(1, Company.NewFullTimeEmployee(1, "Alice", 60000))
	//company.AddEmployee(2, Company.NewPartTimeEmployee(2, "Bob", 80))
	//company.AddEmployee(3, Company.NewFullTimeEmployee(3, "Charlie", 75000))
	//
	//fmt.Println("Company Employees:")
	//for _, details := range company.ListEmployees() {
	//	fmt.Println(details)
	//}

	/*
		------------- Bank--------------------
	*/

	reader := bufio.NewReader(os.Stdin)

	account := Bank.BankAccount{1, "Anuar", 1000}

	var transactions []string

	for {
		fmt.Println("===== Bank Menu =====")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Show Transactions")
		fmt.Println("5. Exit")
		fmt.Println("Choose option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("Enter Deposit amount: ")
			amount := readFloat(reader)

			err := account.Deposit(amount)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				transactions = append(transactions, fmt.Sprintf(": %2.f$ to Account %s deposited successfully!", amount, account.Name))
			}
		case "2":
			fmt.Println("Enter Withdraw amount: ")
			amount := readFloat(reader)

			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				transactions = append(transactions, fmt.Sprintf(": %2.f $ from Account %s withdrawd successfully!", amount, account.Name))
			}

		case "3":
			fmt.Println("Current Balance: \n", account.GetBalance(), "$")

		case "4":
			if len(transactions) == 0 {
				fmt.Println("No transactions found!")
			}
			if len(transactions) > 0 {
				fmt.Println("Transactions:")
				for i, t := range transactions {
					fmt.Println("%d. %s\n", i+1, t)
				}
			}
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
		fmt.Println("")
	}
}

func readFloat(reader *bufio.Reader) float64 {
	text, _ := reader.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)
	return value
}
