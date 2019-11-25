package main

import (
	"encoding/base64"
	"fmt"
	"github.com/alekseymerzlyakov/wuhosted/requests"
	"github.com/alekseymerzlyakov/wuhosted/routes"
	"net/http"
)


var Access_Token string
var authorize_code string
var URLq string = "https://portal.kenya.uat.wuamerigo.com/admin/api/1/oauth2/token"
var urlAuthorize string = "https://portal.kenya.uat.wuamerigo.com/admin/api/1/customers/authorize"


type ViewData struct{
	Available bool
}

	var ClientID string = "KE932150"
	var Secret_Key string = "OAjP#hhV67FLPv_K3SXSIPbh"
	var Aus string = ClientID + ":" + Secret_Key

	//Base64
func Encoded() string {
	encoded := base64.StdEncoding.EncodeToString([]byte(Aus))
	fmt.Println("encoded      " + encoded)
	return encoded
}



func main() {
	//Token
	Access_Token = requests.ReqGetAccess_Token(URLq, Encoded())
	fmt.Println("Access_Token         ----------->>>>>>>>         " + Access_Token)


	//Authorize
	authorize_code = requests.Authorize(urlAuthorize, Access_Token)
	fmt.Println("\n authorize_code         ----------->>>>>>>>         " + authorize_code)









	//
	http.HandleFunc("/", routes.IndexHandler)

	http.ListenAndServe(":8080", nil)
}


