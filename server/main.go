package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gjvnq/go-logger"
	"github.com/gjvnq/wedge2/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Log *logger.Logger

func main() {
	var err error

	// Set Logger
	Log, err = logger.New("main", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	wedge.Log = Log

	// Connect Database
	wedge.DB, err = sql.Open("mysql", GetDBStringURI())
	if err != nil {
		Log.FatalF("Failed to open database %s", err)
	}

	// Listen for connections
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", Books).Methods("GET")
	router.HandleFunc("/auth", Auth).Methods("POST")

	handler := CorsMiddleware(router)
	Log.Notice("Now listening...")
	Log.Fatal(http.ListenAndServe(":8081", handler))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
