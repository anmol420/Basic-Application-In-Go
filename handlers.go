package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemp(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemp(w, "about.page.tmpl")
}