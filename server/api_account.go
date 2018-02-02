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

func AccountBalanceHistoric(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Load account
	acc, err := wedge.Accounts.ByID(GetUUID("acc-id", r))
	if err != nil {
		Log.Warning(err)
		SendErrCodeAndLog(w, 500, err)
		return
	}

	// Check auth
	if !acc.BookID.Equal(GetUUID("book-id", r)) {
		SendErrCode(w, 403)
		return
	}

	// Get data
	err = acc.LoadHistoric(GetLDate("from", r), GetLDate("to", r))
	if err != nil {
		Log.Warning(err)
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, acc.Historic)
}

func AccountBalances(w http.ResponseWriter, r *http.Request) {
	var time wedge.LDate

	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Parse date
	if GetString("time", r) == "now-utc" {
		time = wedge.LDateNow()
	} else {
		time = GetLDate("time", r)
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

func GetAccountsCore(time wedge.LDate, tree_flag bool, w http.ResponseWriter, r *http.Request) {
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

	tree := wedge.AccountTree(accounts)

	if tree_flag {
		sendJSONResponse(w, tree)
	} else {
		sendJSONResponse(w, wedge.AccountList(tree))
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
