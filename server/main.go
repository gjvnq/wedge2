package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gjvnq/wedge2/domain"
	"github.com/gorilla/mux"
)

func main() {
	// Connect Database

	// Listen for connections
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/auth", Auth)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	handler := CorsMiddleware(router)
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
