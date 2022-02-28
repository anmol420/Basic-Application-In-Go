package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portnumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting The Application On Port %s", portnumber))
	
	http.ListenAndServe(portnumber, nil)
}