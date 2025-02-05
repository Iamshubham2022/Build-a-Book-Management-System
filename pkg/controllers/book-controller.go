package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/utils"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/models"
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

func GetBookbyId(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId := vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing...")
	}
	bookDetail, _ := models.GetBookById(ID)
	res,_:=json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createbook(w http.ResponseWriter, r *http.Request){
	createbook := &models.Book{}
	utils.ParseBody(r,createbook)
	b:=createbook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err:= strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while deleting ...")
	}
	deletebook := models.DeleteBookById(ID)
	res,_:=json.Marshal(deletebook)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func Updatabook(w http.ResponseWriter,r http.Request){
	
}
