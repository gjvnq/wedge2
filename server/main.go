package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gjvnq/go-logger"
	uuid "github.com/gjvnq/go.uuid"
	"github.com/gjvnq/wedge2/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Log *logger.Logger
var JWTKey []byte
var Config ConfigS

func main() {
	var err error

	Log, err = logger.New("main", 1, os.Stdout)
	Config = LoadConfigFile()

	// Set Logger
	if Config.DevMode == false {
		Log.Levels["INFO"] = false
		Log.Levels["DEBUG"] = false
	}

	if Config.DevMode {
		Log.Notice("Using development mode")
	} else {
		Log.Notice("Using production mode")
	}

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
	if Config.DevMode {
		JWTKey = make([]byte, 16)
	}

	// Connect Database
	wedge.DB, err = sql.Open("mysql", Config.MySQL)
	if err != nil {
		Log.FatalF("Failed to open database %s", err)
	}
	wedge.PrepareStatments()

	// Listen for connections
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", BooksList).Methods("GET")
	router.HandleFunc("/books/{book-id}/assets", AssetsList).Methods("GET")
	router.HandleFunc("/books/{book-id}/assets", AssetsPut).Methods("PUT")
	router.HandleFunc("/books/{book-id}/accounts", AccountList).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts/balances/at/{time}", AccountBalances).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts-tree", AccountTree).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts", AccountSet).Methods("PUT")
	router.HandleFunc("/books/{book-id}/accounts/{acc-id}/balance/{from}/{to}", AccountBalanceHistoric).Methods("GET")
	router.HandleFunc("/books/{book-id}/accounts/{acc-id}/movements", MovementsAccountGet).Methods("GET")
	router.HandleFunc("/books/{book-id}/transactions", TransactionSet).Methods("PUT")
	router.HandleFunc("/books/{book-id}/transactions", TransactionList).Methods("GET")
	router.HandleFunc("/books/{book-id}/transactions/{tr-id}", TransactionGet).Methods("GET")
	router.HandleFunc("/books/{book-id}/transactions/{tr-id}", TransactionRm).Methods("DELETE")
	router.HandleFunc("/auth", Auth).Methods("POST")
	router.HandleFunc("/auth/test", AuthTest).Methods("POST")

	handler := CorsMiddleware(router)
	Log.Notice("Now listening... on " + Config.ListenOn + ":" + Config.Port)
	Log.Fatal(http.ListenAndServe(Config.ListenOn+":"+Config.Port, handler))
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

func GetString(key string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[key]
}

func GetLDate(key string, r *http.Request) wedge.LDate {
	ldate := wedge.LDate{}
	vars := mux.Vars(r)
	err := ldate.UnmarshalJSON([]byte(vars[key]))
	if err != nil {
		return wedge.LDate{}
	}
	return ldate
}

func SendErrCode(w http.ResponseWriter, code int) {
	switch code {
	case 200:
		http.Error(w, "200 OK", 200)
	case 201:
		http.Error(w, "201 Created", 201)
	case 202:
		http.Error(w, "202 Accepted", 202)
	case 400:
		http.Error(w, "400 Bad Request", 400)
	case 401:
		http.Error(w, "401 Unauthorized Error", 401)
	case 404:
		http.Error(w, "404 Not Found Error", 404)
	case 403:
		http.Error(w, "403 Forbidden Error", 403)
	case 409:
		http.Error(w, "409 Conflict", 409)
	case 500:
		http.Error(w, "500 Internal Server Error", 500)
	default:
		Log.Warning("Unknown http error code:", code)
		http.Error(w, "", code)
	}
}

func SendErrCodeAndLog(w http.ResponseWriter, code int, err interface{}) {
	Log.WarningNF(1, "Sending %d HTTP Error Code due to: %v", code, err)
	SendErrCode(w, code)
}

func Is404(err error) bool {
	return err == sql.ErrNoRows
}

func IsDuplicate(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "duplicate entry")
}
