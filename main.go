package main 

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// r.HandleFunc("/books/{title}", handlers.getBooks)	

	http.ListenAndServe(":3333", r)


	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	queryParams := r.URL.Query()

	// 	id := queryParams.Get("id")

	// 	fmt.Fprintf(w, "Hello, you've requested ID: %s\n", id)
	// 	fmt.Fprintln(w, "Hello, you've requested ID: %s\n", id)

	// })
	// http.ListenAndServe(":3333", nil)
}