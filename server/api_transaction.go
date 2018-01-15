package main

import (
	"encoding/json"
	"net/http"

	"github.com/gjvnq/wedge2/domain"
)

func TransactionList(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Transactions
	transactions, err := wedge.Transactions.InBook(GetBookId(r))
	if err != nil {
		SendErrCodeAndLog(w, 500, err)
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
	Log.Debug(transaction)

	// Check auth
	if IsAuthInvalid3(w, r, &transaction.BookID) {
		return
	}

	// Put on database
	err = wedge.Transactions.Set(&transaction)
	if err != nil {
		if IsDuplicate(err) {
			SendErrCodeAndLog(w, 409, err)
			return
		}
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, transaction)
}

func TransactionGet(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Get Transaction
	tr, err, not_found := wedge.Transactions.GetByID(GetUUID("tr-id", r))
	if not_found {
		SendErrCodeAndLog(w, 404, err)
		return
	}
	if err != nil {
		Log.Warning(GetUUID("tr-id", r), err)
		SendErrCodeAndLog(w, 500, err)
		return
	}

	// Check auth
	if !tr.BookID.Equal(GetBookId(r)) {
		SendErrCodeAndLog(w, 403, err)
		return
	}

	sendJSONResponse(w, tr)
}

func TransactionRm(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Get Transaction
	tr, err, not_found := wedge.Transactions.GetByID(GetUUID("tr-id", r))
	Log.Debug(tr, err, not_found)
	if not_found {
		SendErrCode(w, 404)
		return
	} else if err != nil {
		SendErrCodeAndLog(w, 500, err)
		return
	}

	// Check auth
	if !tr.BookID.Equal(GetBookId(r)) {
		SendErrCode(w, 403)
		return
	}

	// Actually delete
	err = wedge.Transactions.RmByID(tr.ID)
	if err != nil {
		SendErrCodeAndLog(w, 500, err)
		return
	}

	SendErrCode(w, 200)
}
