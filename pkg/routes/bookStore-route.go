package routes

import (
	"github.com/gorilla/mux"
	// "/Users/apple1/Desktop/GolangProject/Build-a-Book-Management-System/pkg/controllers"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")

}
