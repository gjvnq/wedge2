package main

import (
	"encoding/json"
	"net/http"

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
		SendErrCodeAndLog(w, 500, err)
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
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, wedge.AccountTree(accounts))
}

func AccountSet(w http.ResponseWriter, r *http.Request) {
	// Parse data
	account := wedge.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		SendErrCodeAndLog(w, 400, err)
		return
	}

	// Check auth
	if IsAuthInvalid3(w, r, &account.BookID) {
		return
	}

	// Put on database
	err = wedge.Accounts_Set(&account)
	if err != nil {
		if IsDuplicate(err) {
			SendErrCodeAndLog(w, 409, err)
			return
		}
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, account)
}
