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

	GetAccountsCore(wedge.LDateNow(), false, w, r)
}

func AccountBalances(w http.ResponseWriter, r *http.Request) {
	var err error
	var time wedge.LDate

	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Parse date
	if GetString("time", r) == "now-utc" {
		time = wedge.LDateNow()
	} else {
		time, err = GetLDate("time", r, w)
		if err != nil {
			return
		}
	}

	GetAccountsCore(time, false, w, r)
}

func AccountTree(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	GetAccountsCore(wedge.LDateNow(), true, w, r)
}

func GetAccountsCore(time wedge.LDate, tree bool, w http.ResponseWriter, r *http.Request) {
	// List Accounts
	accounts, err := wedge.Accounts.InBook(GetBookId(r))
	if err != nil {
		SendErrCodeAndLog(w, 500, err)
		return
	}
	for i := 0; i < len(accounts); i++ {
		err = accounts[i].LoadBalanceAt(time)
		if err != nil {
			SendErrCodeAndLog(w, 500, err)
			return
		}
	}

	if tree {
		sendJSONResponse(w, wedge.AccountTree(accounts))
	} else {
		sendJSONResponse(w, accounts)
	}
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
	err = wedge.Accounts.Set(&account)
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
