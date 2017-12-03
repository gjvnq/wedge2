package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gjvnq/go-logger"
	uuid "github.com/gjvnq/go.uuid"
	"github.com/gjvnq/wedge2/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Log *logger.Logger
var JWTKey []byte

func main() {
	var err error

	// Set Logger
	Log, err = logger.New("main", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	wedge.Log = Log

	// Generate secure key for JWT
	JWTKey = make([]byte, 16)
	_, err = rand.Read(JWTKey)
	if err != nil {
		Log.Fatal("Failed to generate secure key:", err)
		return
	}
	// Hack for dev
	JWTKey = make([]byte, 16)

	// Connect Database
	wedge.DB, err = sql.Open("mysql", GetDBStringURI())
	if err != nil {
		Log.FatalF("Failed to open database %s", err)
	}

	// Listen for connections
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", BooksList).Methods("GET")
	router.HandleFunc("/books/{book-id}/assets", AssetsList).Methods("GET")
	router.HandleFunc("/books/{book-id}/assets", AssetsPut).Methods("PUT")
	router.HandleFunc("/books/{book-id}/accounts", AccountList).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts-tree", AccountTree).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts", AccountSet).Methods("PUT")
	router.HandleFunc("/auth", Auth).Methods("POST")
	router.HandleFunc("/auth/test", AuthTest).Methods("POST")

	handler := CorsMiddleware(router)
	Log.Notice("Now listening...")
	Log.Fatal(http.ListenAndServe(":8081", handler))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
		}
	})
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		Log.WarningNF(1, "Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		Log.WarningNF(1, "Failed to write the response body: %v", err)
		return
	}
}

func sendResponse(w http.ResponseWriter, mime string, data []byte) {
	var err error
	if mime != "" {
		w.Header().Set("Content-Type", mime)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		Log.WarningNF(1, "Failed to write the response body: %v", err)
		return
	}
}

func GetBookId(r *http.Request) uuid.UUID {
	return GetUUID("book-id", r)
}

func GetUUID(key string, r *http.Request) uuid.UUID {
	vars := mux.Vars(r)
	return uuid.FromStringOrNil(vars[key])
}
