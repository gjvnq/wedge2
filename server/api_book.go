package main

import (
	"net/http"

	"github.com/gjvnq/wedge2/domain"
)

func BooksList(w http.ResponseWriter, r *http.Request) {
	// Load books
	books, err := wedge.Books_All(true)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, books)
}
