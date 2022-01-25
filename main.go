package main

import (
	"fmt"
	"net/http"
)

const portnumber = ":3000"

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting The Application On Port %s", portnumber))
	
	http.ListenAndServe(portnumber, nil)
}