package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

type AuthReq struct {
	BookId   uuid.UUID `json:"bookId"`
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

	// Debug
	fmt.Println(auth_req)

	// Check with DB
}

// w.WriteHeader(http.StatusForbidden)
