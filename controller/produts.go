package controller

import (
	"html/template"
	"net/http"
	"ze-is-store/db"
)

var pages = template.Must(template.ParseGlob("../templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProduts := db.GetAllProduts()
	pages.ExecuteTemplate(w, "Index", allProduts)
}

func New(w http.ResponseWriter, r *http.Request) {
	pages.ExecuteTemplate(w, "New", nil)
}
