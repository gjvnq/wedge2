package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gjvnq/wedge2/domain"
)

func AccountList(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Accounts
	accounts, err := wedge.Accounts_InBook(GetBookId(r))
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, accounts)
}

func AccountTree(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Accounts
	accounts, err := wedge.Accounts_InBook(GetBookId(r))
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, wedge.AccountTree(accounts))
}

func AccountSet(w http.ResponseWriter, r *http.Request) {
	// Parse data
	account := wedge.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Check auth
	if IsAuthInvalid3(w, r, &account.BookID) {
		return
	}

	// Put on database
	err = wedge.Accounts_Set(&account)
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

	sendJSONResponse(w, account)
}
