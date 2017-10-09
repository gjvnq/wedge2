package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gjvnq/wedge2/domain"
    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	trn := wedge.Transaction{}
    a, b := wedge.Div(101, 3)
    fmt.Fprintf(w, "%+v %+v\n", a, b)
    a, b = wedge.NSplit(101, 3)
    fmt.Fprintf(w, "%+v %+v\n", a, b)
    fmt.Fprintf(w, "%+v\n", trn)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}