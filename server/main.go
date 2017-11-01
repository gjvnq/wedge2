package main

import (
	"database/sql"
	"fmt"
	"log"
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
	router.HandleFunc("/", Index)
	router.HandleFunc("/auth", Auth)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	handler := CorsMiddleware(router)
	Log.Notice("Now listening...")
	log.Fatal(http.ListenAndServe(":8081", handler))
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

func Index(w http.ResponseWriter, r *http.Request) {
	trn := wedge.Transaction{}
	a, b := wedge.Div(101, 3)
	fmt.Fprintf(w, "%+v %+v\n", a, b)
	a, b = wedge.NSplit(101, 3)
	fmt.Fprintf(w, "%+v %+v\n", a, b)
	fmt.Fprintf(w, "%+v\n", trn)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
