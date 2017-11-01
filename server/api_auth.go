package main

import (
	"encoding/json"
	"net/http"

	"github.com/gjvnq/go.uuid"
	"github.com/gjvnq/wedge2/domain"
)

type AuthReq struct {
	BookID   uuid.UUID `json:"book_id"`
	Password string    `json:"password"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	// Load request parameters
	auth_req := AuthReq{}
	err := json.NewDecoder(r.Body).Decode(&auth_req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if auth_req.BookID.IsNil() {
		http.Error(w, "ID must not be null", 404)
		return
	}

	// Load book
	book, err, not_found := wedge.Books_GetByID(auth_req.BookID)
	if not_found {
		http.Error(w, err.Error(), 404)
		return
	}
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	// Check password
	if book.CheckPassword(auth_req.Password) == false {
		Log.InfoF("Wrong password for %s", auth_req.BookID.String())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	Log.InfoF("Right password for %s", auth_req.BookID.String())

	w.WriteHeader(http.StatusOK)
}
