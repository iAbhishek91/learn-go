package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//MODEL
type Author struct {
	FirstName string `json:"firstName`
	Lastname  string `json:"lastname"`
}
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// DATABASE
var books []Book

// ROUTE HANDLERS
// route handler should have these two properties
func getBooks(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	// response with all the books
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get parameter from request

	// loop thorough all the books and send the book requested for
	// This can be select query to the database
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return // stop the iteration and return from the function
		}
	}

	// IF Book not found, return empty structure
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// read the request body and save it in book object
	_ = json.NewDecoder(r.Body).Decode(&book)

	// generate a random number and convert that to string
	book.ID = strconv.Itoa(rand.Intn(1000000))

	// append the book to books database
	books = append(books, book)

	// response with the book created
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// read the request body
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			var book Book

			// read the request body and save it in book object
			_ = json.NewDecoder(r.Body).Decode(&book)

			// generate a random number and convert that to string
			book.ID = params["id"]

			// append the book to books database
			books = append(books, book)

			// response with the book created
			json.NewEncoder(w).Encode(book)

			return
		}
	}

	// response with all the books
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// read the request body
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	// response with all the books
	json.NewEncoder(w).Encode(books)
}

// MAIN
func main() {
	// Initialize router
	r := mux.NewRouter()

	// MOCK DATA
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "123123",
			Title: "Book one",
			Author: &Author{
				FirstName: "Rob",
				Lastname:  "Dickson",
			},
		},
	)
	books = append(
		books,
		Book{
			ID:    "2",
			Isbn:  "123124",
			Title: "Book two",
			Author: &Author{
				FirstName: "Tom",
				Lastname:  "Dickson",
			},
		},
	)
	books = append(
		books,
		Book{
			ID:    "3",
			Isbn:  "123125",
			Title: "Book three",
			Author: &Author{
				FirstName: "Bob",
				Lastname:  "Dickson",
			},
		},
	)
	// API: localhost:3002/api/books | GET
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	// API: localhost:3002/api/books/1 | GET
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// API: localhost:3002/api/books | HEADER: Content-Type: application/type; | POST
	// BODY
	/*
			{
		    "isbn": "34343",
		    "title": "Book Four",
		    "author": {
		        "firstName": "Dom",
		        "lastname": "Jackson"
		    }
		}
	*/
	r.HandleFunc("/api/books", createBook).Methods("POST")
	// API: localhost:3002/api/books/1 | PUT
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// API: localhost:3002/api/books/1 | DELETE
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// API: localhost:3002/api/books/1 | HEADER: Content-Type: application/type; | PUT
	// BODY
	/*
			{
		    "isbn": "34343",
		    "title": "Book Four",
		    "author": {
		        "firstName": "Dom",
		        "lastname": "Jackson"
		    }
		}
	*/
	fmt.Println("Server starting at 3002...")
	log.Fatal(http.ListenAndServe(":3002", r))
}
