package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gjvnq/wedge2/domain"
)

func TransactionList(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Transactions
	transactions, err := wedge.Transactions_InBook(GetBookId(r))
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, transactions)
}

func TransactionSet(w http.ResponseWriter, r *http.Request) {
	// Parse data
	transaction := wedge.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Check auth
	if IsAuthInvalid3(w, r, &transaction.BookID) {
		return
	}

	// Put on database
	err = wedge.Transactions_Set(&transaction)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err != nil {
				http.Error(w, "duplicate entry", 409)
				return
			}
		}
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, transaction)
}
