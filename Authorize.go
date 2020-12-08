package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Respon struct {
	Authoriz    string `json:"authorization_code"`
	Action_info string `json:"action_info"`
}

func Authorize_Request() []byte {
	fmt.Println("open func Authorize_Request()")

	data := []byte(`{
  "customer_id": "` + GetString("customer_id") + `",
  "profile": {
	"first_name": "` + GetString("first_name") + `",
	"last_name": "` + GetString("last_name") + `",
	"street1": "` + GetString("street1") + `",
	"city": "` + GetString("city") + `",
	"zip_code": "` + GetString("zip_code") + `",
	"country": "` + GetString("country") + `",
	"phone_country_code": "` + GetString("phone_country_code") + `",
	"phone_number": "` + GetString("phone_number") + `",
	"birth": "` + GetString("birth") + `",
	"country_of_birth": "` + GetString("country_of_birth") + `",
	"id_type": "` + GetString("id_type") + `",
	"id_number": "` + GetString("id_number") + `",
	"id_issuer": "` + GetString("id_issuer") + `",
	"id_expires": ` + GetString("id_expires") + `,
	"id_expiration": "` + GetString("id_expiration") + `"
  },
  "bank_accounts": [
	{
  	"account_number": "` + GetString("account_number") + `",
	"routing_number": "22345612346",
  	"name": "` + GetString("name") + `",
	"type": "savings",
  	"currency": "` + Currency + `",
  	"balance": "` + GetString("balance") + `",
	"limit": "10000"
	 }
  ]
}`)

	return data
}

// Post Request
func Authorize(url, Access_Token, X_Time, data string) string {

	fmt.Println("Authorize url -> ", url)
	fmt.Println("Authorize X-Signature -> ", Access_Token)
	fmt.Println("Authorize X_Time -> ", X_Time)
	fmt.Println("Authorize X-Id -> ", PublicKey)
	fmt.Println("Authorize data -> ", string(data))

	log.Println("Authorize url -> ", url)
	log.Println("Authorize X-Signature -> ", Access_Token)
	log.Println("Authorize X_Time -> ", X_Time)
	log.Println("Authorize X-Id -> ", PublicKey)
	log.Println("Authorize data -> ", string(data))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal("Error reading requests. ", err)
	}
	//"balance": "` + GetString("balance2") + `"
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", Access_Token)
	req.Header.Set("X-Time", X_Time)
	req.Header.Set("X-Id", PublicKey)

	// Create and Add cookie to request
	//cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	//req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10000}

	// Validate cookie and headers are attached
	//fmt.Println(req.Cookies())
	//fmt.Println(req.Header)
	//fmt.Println(req.Body)

	fmt.Println("req -> ", req)
	log.Println("req -> ", req)
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:   ---- >    ", resp.Status)
	fmt.Println("response Headers:   ---- >    ", resp.Header)
	log.Println("response Status:   ---- >    ", resp.Status)
	log.Println("response Headers:   ---- >    ", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	log.Println("response body:   ---- >    ", string(body))
	//log.Info("%s\n", body)
	//Parse JSON

	//var respon  Respon;
	textBytes := []byte(body)
	respon := Respon{}
	jsonErr := json.Unmarshal(textBytes, &respon)
	if jsonErr != nil {
		//log.Fatal(jsonErr)
	}

	return respon.Authoriz
}
