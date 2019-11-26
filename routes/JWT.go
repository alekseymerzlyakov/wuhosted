package routes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	_ "encoding/json"
	"fmt"
	"github.com/alekseymerzlyakov/wuhosted/requests"
)

var HashSt string
var ClientID string = "KE932150"
var Secret_Key string = "OAjP#hhV67FLPv_K3SXSIPbh"
var Aus string = ClientID + ":" + Secret_Key

var Access_Token string
var authorize_code string
var URLq string = "https://portal.kenya.uat.wuamerigo.com/admin/api/1/oauth2/token"
var urlAuthorize string = "https://portal.kenya.uat.wuamerigo.com/admin/api/1/customers/authorize"

type HashTag struct {
	HashT string
}

//Base64
func Encoded(base string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(base))
	fmt.Println("encoded      " + encoded)
	return encoded
}






func JWT ()  {
	//Token
	Access_Token = requests.ReqGetAccess_Token(URLq, Encoded(Aus))
	fmt.Println("\n Access_Token         ----------->>>>>>>>         " + Access_Token)


	//Authorize
	authorize_code = requests.Authorize(urlAuthorize, Access_Token)
	fmt.Println("\n authorize_code         ----------->>>>>>>>         " + authorize_code)

	var header string = "{\"typ\":\"JWT\",\"alg\":\"HS256\"}";
	var payload string = "{\"authorization_code\":\""+ authorize_code +"\"}";
	var JWT_header_payload string = Encoded(header) + "." + Encoded(payload)
	var secretKey string = "6xrw38YLfIc4e!J2dw5a1OeE8JZOrBCS";

	HashSt = Hash(JWT_header_payload, secretKey)
	fmt.Printf("\n hash      --------------->>>        " + HashSt)
	fmt.Printf("\n JWT_header_payload      --------------->>>        " + JWT_header_payload)
	fmt.Printf("\n payload      --------------->>>        " + payload)



}




// Hash generates a Hmac256 hash of a string using a secret
func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	var jw = src + "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	return jw
}

// isValidHash validates a hash againt a value
func isValidHash(value string, hash string, secret string) bool {
	return hash == Hash(value, secret)
}




