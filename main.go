package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	http.ListenAndServe(":3000", r)
}
