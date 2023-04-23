package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafliputraa/learn-mysql-bookstore-go/pkg/models"
	"github.com/rafliputraa/learn-mysql-bookstore-go/pkg/utils"
)

var NewBook models.Book

func GetBook(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	newBooks := models.GetAllBooks()
	responseData, _ := json.Marshal(newBooks)
	responseWriter.Header().Set("Content-Type", "pkglication/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseData)
}

func GetBookById(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	responseData, _ := json.Marshal(bookDetails)
	responseWriter.Header().Set("Content-Type", "pkglication/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseData)
}

func CreateBook(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	createBookModel := &models.Book{}
	utils.ParseBody(httpRequest, createBookModel)
	createBook := createBookModel.CreateBook()
	responseData, _ := json.Marshal(createBook)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseData)
}

func DeleteBook(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails := models.DeleteBook(id)
	responseData, _ := json.Marshal(bookDetails)
	responseWriter.Header().Set("Content-Type", "pkglication/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseData)
}

func UpdateBook(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(httpRequest, updateBook)

	vars := mux.Vars(httpRequest)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails, db := models.GetBookById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	responseData, _ := json.Marshal(bookDetails)
	responseWriter.Header().Set("Content-Type", "pkglication/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(responseData)

}
