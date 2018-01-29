package main

import (
	"net/http"

	wedge "github.com/gjvnq/wedge2/domain"
)

func MovementsAccountGet(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// Load movements
	movs, err := wedge.Movements.InAccountAndBook(GetUUID("acc-id", r), GetBookId(r))
	if err != nil {
		Log.Warning(err)
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, movs)
}
