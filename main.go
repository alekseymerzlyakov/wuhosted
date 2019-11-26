package main

import (
	"github.com/alekseymerzlyakov/wuhosted/routes"
	"net/http"
)

type ViewData struct{
	Available bool
}




func main() {

	routes.JWT()

	//
	http.HandleFunc("/", routes.IndexHandler)

	http.ListenAndServe(":8080", nil)
}


