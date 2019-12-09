package routes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
)

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
	//Access_Token = ReqGetAccess_Token(URLq, Encoded(Aus))

	//Authorize
	authorize_code = Authorize(UrlAuthorize, Access_Token)

	var header string = "{\"typ\":\"JWT\",\"alg\":\"HS256\"}"
	var payload string = "{\"authorization_code\":\""+ authorize_code +"\"}"
	JWT_header_payload  = Encoded(header) + "." + Encoded(payload)
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



