package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harshad-salunke2002/books_api/models"
	"github.com/harshad-salunke2002/books_api/response"
)

// endpoints "/book/books"  for get all books
func (db *BookDB) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := db.ReadAllBook()

	if err != nil {
		response.ResponseWithError(w, http.StatusNotFound, "Error while Reading Record")
		return
	}

	response.ResponseWithJson(w, http.StatusFound, books)

}

// endpoints "book/addbook"  for adding books
func (db *BookDB) AddBook(w http.ResponseWriter, r *http.Request) {

	//parsing body
	var newBook models.Books
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, fmt.Sprint("Error while decoding json : ", err.Error()))
		return
	}

	//validating book

	isBookValid, msg := ValidateBook(&newBook)

	if !isBookValid {
		response.ResponseWithError(w, http.StatusBadRequest, fmt.Sprint("Missing Data: ", msg))
		return
	}

	//inserting book to db
	isDataInserted, err := db.InsertBook(&newBook)
	fmt.Println(newBook)

	if !isDataInserted {
		response.ResponseWithError(w, http.StatusBadRequest, fmt.Sprint("Error while Inserting data : ", err.Error()))
		return
	}

	response.ResponseWithJson(w, http.StatusOK, models.SuccessResponse{
		Msg:     "Book Added successfully",
		Success: true,
	})

}

// endpoints "/book/getbook/2"  for get book by its ID
func (db *BookDB) GetBookByID(w http.ResponseWriter, r *http.Request) {
	//parsing id
	params := mux.Vars(r)
	bookIdStr := params["bookId"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, "Error while Parsing Book ID")
		return
	}

	//reading book from db
	book, err := db.ReadBook(bookId)
	if err != nil {
		response.ResponseWithError(w, http.StatusNotFound, fmt.Sprint("Recode Not Found : ", err.Error()))
		return
	}

	response.ResponseWithJson(w, http.StatusFound, book)

}

// endpoints "/book/update/2"  for updating book
func (db *BookDB) UpdateBook(w http.ResponseWriter, r *http.Request) {

	//parsing id
	params := mux.Vars(r)
	bookIdStr := params["bookId"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, "Error while Parsing Book ID")
		return
	}

	//parsing body
	var newBook models.Books

	error := json.NewDecoder(r.Body).Decode(&newBook)

	if error != nil {
		response.ResponseWithError(w, http.StatusBadRequest, fmt.Sprint("Error while decoding json : ", error.Error()))
		return
	}

	fmt.Println("before validating")

	//validating book
	isBookValid, msg := ValidateBook(&newBook)

	if !isBookValid {
		response.ResponseWithError(w, http.StatusBadRequest, fmt.Sprint("Missing Data: ", msg))
		return
	}

	fmt.Println("after validating")

	//updating error
	up_err := db.UpdateBookDB(bookId, &newBook)
	fmt.Println("after update")

	if up_err != nil {
		response.ResponseWithError(w, http.StatusNotFound, fmt.Sprint("Recode Not Found : ", up_err.Error()))
		return
	}

	response.ResponseWithJson(w, http.StatusOK, models.SuccessResponse{
		Msg:     "Book Updated successfully",
		Success: true,
	})

}

func (db *BookDB) DeleteBook(w http.ResponseWriter, r *http.Request) {

	//parsing id
	params := mux.Vars(r)
	bookIdStr := params["bookId"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		response.ResponseWithError(w, http.StatusBadRequest, "Error while Parsing Book ID")
		return
	}

	//deleting book from db
	dele_err := db.DeleteBookbyId(bookId)
	if dele_err != nil {
		response.ResponseWithError(w, http.StatusNotFound, dele_err.Error())
		return
	}
	response.ResponseWithJson(w, http.StatusOK, models.SuccessResponse{
		Msg:     "Book Deleted successfully",
		Success: true,
	})

}

func ValidateBook(book *models.Books) (bool, string) {
	if book.Name == "" {
		return false, "Missing Name Field"
	}

	if book.Writer == "" {
		return false, "Missing Writer Field"
	}

	return true, ""
}
