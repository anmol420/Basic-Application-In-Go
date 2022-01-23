package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portnumber = ":3000"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemp(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemp(w, "about.page.tmpl")
}

func renderTemp(w http.ResponseWriter, tmpl string) {
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error occured", err)
		return
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting The Application On Port %s", portnumber))
	
	http.ListenAndServe(portnumber, nil)
}