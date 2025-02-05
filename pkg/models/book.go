package models

import (
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectORMToDataBase()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBook() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book,*gorm.DB){
	var getBook Book
	db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}


// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	result := db.Where("ID=?", Id).First(&getBook)  // `First` will return the first matched record or an error
// 	if result.Error != nil {
// 		return nil, result // Return nil and the error if no result is found
// 	}
// 	return &getBook, result
// }

/*
var getBook Book

A new variable getBook of type Book is created,
 which will hold the result of the query.
db.Where("ID=?", Id).Find(&getBook)

Where("ID=?", Id) filters the database query to select the record where the ID matches the given Id.
.Find(&getBook) executes the query and populates the getBook variable with the result.
The Find function will execute a SELECT query to fetch the book based on the provided ID.
return &getBook, db

The method returns a pointer to the getBook object and the db object.
The returned pointer allows you to access the Book struct and its fields,
 while the db object can be used for further database operations if needed.
*/

func DeleteBookById(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}



// func DeleteBookById(Id int64) (*Book, error) {
// 	var book Book
// 	result := db.Where("ID=?", Id).First(&book)  // Check if the book exists
// 	if result.Error != nil {
// 		return nil, result.Error  // Return nil and the error if the book doesn't exist
// 	}

// 	// Proceed with deleting the book
// 	db.Delete(&book)
// 	return &book, nil  // Return the deleted book and no error
// }

/*
var book Book

A new variable book of type Book is declared, which will hold the record that is to be deleted.
db.Where("ID=?", Id).Delete(book)

Where("ID=?", Id): Filters the database query to select the record where the ID matches the given Id.
.Delete(book): Executes the delete query to remove the book record from the database.
 This operation will delete the book from the database.
return book

The method returns the book object. However, since the record has been deleted,
 this book will not contain any data from the database (i.e., it will be an empty struct).
*/