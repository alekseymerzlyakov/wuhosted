package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
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

func JWT() {
	fmt.Println("---------->          JWT2")
	request_body_autorize_ := Authorize_Request() //
	request_body_autorize := string(request_body_autorize_)

	request_body_autorize = strings.ReplaceAll(request_body_autorize, " ", "")
	request_body_autorize = strings.ReplaceAll(request_body_autorize, "\r", "")
	request_body_autorize = strings.ReplaceAll(request_body_autorize, "\n", "")

	fmt.Println("Secret      " + SecretKey)
	log.Println("Secret      " + SecretKey)

	fmt.Println("request_body_autorize     ---->        " + request_body_autorize)
	log.Println("request_body_autorize     ---->        " + request_body_autorize)

	X_Time := Time()
	fmt.Println("X_Time      " + X_Time)

	Path := "/api/2/customers/authorize"
	fmt.Println("Path      " + Path)
	log.Println("Path      " + Path)

	Content := Path + request_body_autorize
	combined := X_Time + SecretKey + Content + SecretKey

	generatedSignature := sha_256(combined)

	////Authorize
	authorize_code = Authorize(UrlAuthorize, generatedSignature, X_Time, request_body_autorize)
	fmt.Println("authorize_code ---- >>>   ", authorize_code)
	log.Println("authorize_code ---- >>>   ", authorize_code)

	var header string = "{\"typ\":\"JWT\",\"alg\":\"HS256\"}"
	var payload string = "{\"authorization_code\":\"" + authorize_code + "\"}"
	JWT_header_payload = Encoded(header) + "." + Encoded(payload)
	//var secretKey string = GetString("secretKey")

	HashSt = Hash(JWT_header_payload, SecretKey)
}

// Hash generates a Hmac256 hash of a string using a secret
func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	var jw = src + "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println("JWT      ", jw)
	log.Println("JWT      ", jw)
	return jw
}
