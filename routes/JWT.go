package routes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
	"github.com/magiconair/properties"
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

//Open Properties file
func GetString(key string) string {
	p := properties.MustLoadFile(path() + "/wu.properties", properties.UTF8)
	return p.MustGet(key)
}

var HashSt string
var ClientID string = GetString("ClientID")
var Secret_Key string = GetString("Secret_Key")
var Aus string = ClientID + ":" + Secret_Key
var Access_Token string
var authorize_code string
var URLq string = "https://" + GetString("URL") + "/admin/api/1/oauth2/token"
var urlAuthorize string = "https://" + GetString("URL") + "/admin/api/1/customers/authorize"


type HashTag struct {
	HashT string
}

//Base64
func Encoded(base string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(base))
	log.Println("encoded      " + encoded)
	return encoded
}

func JWT ()  {
	//Token
	Access_Token = ReqGetAccess_Token(URLq, Encoded(Aus))

	//Authorize
	authorize_code = Authorize(urlAuthorize, Access_Token)

	var header string = "{\"typ\":\"JWT\",\"alg\":\"HS256\"}"
	var payload string = "{\"authorization_code\":\""+ authorize_code +"\"}"
	var JWT_header_payload string = Encoded(header) + "." + Encoded(payload)
	var secretKey string = GetString("secretKey")

	HashSt = Hash(JWT_header_payload, secretKey)
}

// Hash generates a Hmac256 hash of a string using a secret
func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	var jw = src + "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	return jw
}



