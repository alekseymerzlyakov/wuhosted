package main

import (
	"github.com/alekseymerzlyakov/wuhosted/routes"
	"net/http"
)

func main() {
	routes.Logrus()
	routes.JWT()

	http.HandleFunc("/", routes.IndexHandler)

	http.ListenAndServe(":8080", nil)
}


