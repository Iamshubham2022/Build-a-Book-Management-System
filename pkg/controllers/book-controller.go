package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/models"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newbook := models.GetAllBook()
	res, _ := json.Marshal(newbook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing...")
	}
	bookDetail, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createbook := &models.Book{}
	utils.ParseBody(r, createbook)
	b := createbook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while deleting ...")
	}
	deletebook := models.DeleteBookById(ID)
	res, _ := json.Marshal(deletebook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	// r in json formate but golang understand struct fromate for creating or updating somthing
	var updatabook = &models.Book{}
	utils.ParseBody(r, updatabook) // now r is converting struct formate in updatebook
	// which id we will want to update
	vars := mux.Vars(r)
	bookId := vars["bookId"] // this is in string formate so convert it in int from

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing.... ")
	}
	// first of all get that book and update  that book
	bookDetails, db := models.GetBookById(ID)

	if updatabook.Name != "" {
		bookDetails.Name = updatabook.Name
	}
	if updatabook.Author != "" {
		bookDetails.Author = updatabook.Author
	}
	if updatabook.Publication != "" {
		bookDetails.Publication = updatabook.Publication
	}
	// after updation is done we need to save the book
	db.Save(&bookDetails)
	// then return the updated book  in json format or response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
