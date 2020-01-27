package main

import (
	"github.com/alekseymerzlyakov/wuhosted/routes"
	"net/http"
)

func main() {
	//routes.Logrus()
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/approvesend", routes.IndexApproveSend)
	http.ListenAndServe(":8080", nil)
}
