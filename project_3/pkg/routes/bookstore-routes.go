package routes

import (
	"project_3/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
