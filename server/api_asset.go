package main

import (
	"encoding/json"
	"net/http"

	"github.com/gjvnq/wedge2/domain"
)

func AssetsList(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Assets
	assets, err := wedge.Assets.InBook(GetBookId(r))
	if err != nil {
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, assets)
}

func AssetsPut(w http.ResponseWriter, r *http.Request) {
	// Parse data
	asset := wedge.Asset{}
	err := json.NewDecoder(r.Body).Decode(&asset)
	if err != nil {
		SendErrCodeAndLog(w, 400, err)
		return
	}

	// Check auth
	if IsAuthInvalid3(w, r, &asset.BookID) {
		return
	}

	// Put on database
	err = wedge.Assets.Set(&asset)
	if err != nil {
		if IsDuplicate(err) {
			SendErrCodeAndLog(w, 409, err)
			return
		}
		SendErrCodeAndLog(w, 500, err)
		return
	}

	sendJSONResponse(w, asset)
}
