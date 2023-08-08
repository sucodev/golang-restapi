package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	title := vars["title"]

	fmt.Fprintf(w, "You've request the book: %s ", title)
}