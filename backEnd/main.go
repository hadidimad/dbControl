package main

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

//GetTables :write database tables to client in json
func GetTables(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tables", GetTables)
	http.ListenAndServe(":8080", router)
}
