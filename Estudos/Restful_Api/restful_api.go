package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/funcionarios", FuncIndex)
    router.HandleFunc("/funcionarios/{funcId}", FuncShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func FuncIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Olá funcionario!")
}

func FuncShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    funcId := vars["funcId"]
    fmt.Fprintln(w, "Funcionário:", funcId)
}
