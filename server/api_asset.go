package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gjvnq/wedge2/domain"
)

func AssetsList(w http.ResponseWriter, r *http.Request) {
	// Check auth
	if IsAuthInvalid(w, r) {
		return
	}

	// List Assets
	assets, err := wedge.Assets_InBook(GetBookId(r))
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	sendJSONResponse(w, assets)
}

func AssetsPut(w http.ResponseWriter, r *http.Request) {
	// Parse data
	asset := wedge.Asset{}
	err := json.NewDecoder(r.Body).Decode(&asset)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Check auth
	if IsAuthInvalid3(w, r, &asset.BookID) {
		return
	}

	// Put on database
	err = wedge.Assets_Insert(&asset)
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

	sendJSONResponse(w, asset)
}
