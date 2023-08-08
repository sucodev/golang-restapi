package main 

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sucodev/golang-restapi/handlers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/books/{title}", handlers.GetBooks).Methods("GET");	

	http.ListenAndServe(":3333", r)


	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	queryParams := r.URL.Query()

	// 	id := queryParams.Get("id")

	// 	fmt.Fprintf(w, "Hello, you've requested ID: %s\n", id)
	// 	fmt.Fprintln(w, "Hello, you've requested ID: %s\n", id)

	// })
	// http.ListenAndServe(":3333", nil)
}