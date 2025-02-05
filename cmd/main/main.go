package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Iamshubham2022/Build-a-Book-Management-System/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("server successfully started...")
}
