package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gjvnq/go.uuid"
	"github.com/gjvnq/wedge2/domain"
	"github.com/gorilla/mux"
)

type AuthReq struct {
	BookID   uuid.UUID `json:"book_id"`
	Password string    `json:"password"`
}

const DEFAULT_TOKEN_LIFE = 600 * time.Minute

func Auth(w http.ResponseWriter, r *http.Request) {
	// Load request parameters
	auth_req := AuthReq{}
	err := json.NewDecoder(r.Body).Decode(&auth_req)
	if err != nil {
		SendErrCodeAndLog(w, 400, err)
		return
	}
	if auth_req.BookID.IsNil() {
		SendErrCodeAndLog(w, 400, "ID must not be null")
		return
	}

	// Load book
	book, err, not_found := wedge.Books.GetByID(auth_req.BookID)
	if not_found {
		SendErrCodeAndLog(w, 404, err)
		return
	}
	if err != nil {
		SendErrCodeAndLog(w, 500, err)
		return
	}

	// Check password
	if book.CheckPassword(auth_req.Password) == false {
		Log.WarningF("Wrong password for %s", auth_req.BookID.String())
		book.SetPassword(auth_req.Password)
		Log.Debug("May be you want to replace the passsword? ", string(book.Password))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	Log.InfoF("Right password for %s", auth_req.BookID.String())

	// Generate JWT
	claims := jws.Claims{}
	claims.SetIssuedAt(time.Now())
	claims.SetSubject(book.ID.String())
	claims.SetExpiration(time.Now().Add(DEFAULT_TOKEN_LIFE))
	token := jws.NewJWT(claims, crypto.SigningMethodHS256)
	b, err := token.Serialize(JWTKey)
	if err != nil {
		Log.Warning("Failed to generate token:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendResponse(w, "", b)
}

func AuthTest(w http.ResponseWriter, r *http.Request) {
	if IsAuthInvalid(w, r) {
		return
	}
	sendResponse(w, "", []byte("ok"))
}

func IsAuthInvalid(w http.ResponseWriter, r *http.Request) bool {
	// Get Book Id
	vars := mux.Vars(r)
	// Parse JWT
	token, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		Log.WarningF("Unparsable JWT: %s", string(r.Header.Get("Authorization")))
		SendErrCodeAndLog(w, 401, err)
		return true
	}
	// Validate JWT
	if token.Validate(JWTKey, crypto.SigningMethodHS256) != nil {
		Log.WarningF("Invalid JWT: %+v", token.Claims())
		SendErrCodeAndLog(w, 401, err)
		return true
	}
	// Check subject
	right_sub := vars["book-id"]
	if sub, _ := token.Claims().Subject(); sub != right_sub {
		Log.WarningF("Wrong subject on JWT. Got %s Expected %s", sub, right_sub)
		SendErrCodeAndLog(w, 401, err)
		return true
	}
	return false
}

func IsAuthInvalid3(w http.ResponseWriter, r *http.Request, book_id *uuid.UUID) bool {
	if IsAuthInvalid(w, r) {
		return true
	}
	vars := mux.Vars(r)
	book_id_str := book_id.String()
	if book_id_str == "" || book_id.IsNil() {
		err := book_id.UnmarshalText([]byte(vars["book-id"]))
		if err != nil {
			SendErrCodeAndLog(w, 400, err)
			return true
		}
		return false
	}
	if book_id_str != vars["book-id"] {
		Log.WarningF("BookId in URL does not match JWT subjetc. URL -> %s JWT -> %s", book_id_str, vars["book-id"])
		w.WriteHeader(http.StatusForbidden)
		return true
	}
	return false
}
