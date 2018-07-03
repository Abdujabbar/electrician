package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	mxRouter := mux.NewRouter()
	mxRouter.HandleFunc("/", serveTemplate)
	mxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8000", mxRouter)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates/layouts", "base.html")
	fp := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
