package routes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/alekseymerzlyakov/wuhosted/requests"
	"github.com/magiconair/properties"
	log "github.com/Sirupsen/logrus"
	"os"
)

//var path string
func path () string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func getString(key string) string {
	p := properties.MustLoadFile(path() + "/wu.properties", properties.UTF8)
	return p.MustGet(key)
}

var HashSt string
var ClientID string = getString("ClientID")
var Secret_Key string = getString("Secret_Key")
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

	log.Error("Not fatal. An error. Won't stop execution")

	//Token
	Access_Token = requests.ReqGetAccess_Token(URLq, Encoded(Aus))
	fmt.Println("\n Aus         ----------->>>>>>>>         " + Aus)
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



