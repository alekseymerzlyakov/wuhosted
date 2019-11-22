package main

import (
	"fmt"

	"github.com/alekseymerzlyakov/wuhosted/routes"
	"net/http"
)


type ViewData struct{
	Available bool
}




func main() {
	fmt.Println("Aleks \n")


	//request.ReqGet()

	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":8181", nil)
}
