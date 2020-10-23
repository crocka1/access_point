package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	var err error

	homeTemplate, err = template.ParseFiles("views/home.gohtml")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	// Assets
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	http.ListenAndServe(":3000", r)
}
