package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Respones struct {
	Authoriz    string `json:"authorization_code"`
	Action_info string `json:"action_info"`
}

// Post Request
func ChangeStatus(IdNum, UrlAuthorize string) string {

	data := []byte(`[{"ID":"` + IdNum + `"}]`)                  // Body
	URL_Admin_full := WuHosted_Portal + "/admin" + UrlAuthorize //
	fmt.Println("URL_Admin_full -> ", URL_Admin_full)

	//
	request_body_autorize_ := data //
	request_body_autorize := string(request_body_autorize_)
	request_body_autorize = strings.Replace(request_body_autorize, "\r", "", -1)
	request_body_autorize = strings.Replace(request_body_autorize, "\n", "", -1)
	request_body_autorize = strings.ReplaceAll(request_body_autorize, " ", "")

	X_Time := Time()

	Content := UrlAuthorize + request_body_autorize
	combined := X_Time + SecretKey + Content + SecretKey

	generatedSignature := sha_256(combined)
	fmt.Println("generatedSignature  ->  " + generatedSignature)
	fmt.Println("X_Time  ->  " + X_Time)
	fmt.Println("PublicKey  ->  " + PublicKey)
	fmt.Println("X_Time + SecretKey + Content + SecretKey  ->  " + combined)
	fmt.Println("Body  ->  " + request_body_autorize)

	//log.Info("request -> " + string(data))
	req, err := http.NewRequest("PUT", URL_Admin_full, bytes.NewBuffer([]byte(request_body_autorize)))
	if err != nil {
		//	log.Fatal("Error reading requests. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Signature", generatedSignature)
	req.Header.Set("X-Time", X_Time)
	req.Header.Set("X-Id", PublicKey)

	// Create and Add cookie to request
	//cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	//req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10000}

	// Validate cookie and headers are attached
	//fmt.Println(req.Cookies())
	//log.Info(req.Header)
	//log.Info(req.Body)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		//	log.Fatal("Error reading response. ", err)
	}

	//log.Print(resp.Body)

	defer resp.Body.Close()

	//log.Info("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//	log.Fatal("Error reading body. ", err)
	}

	//log.Info("%s\n", string(body))
	//Parse JSON

	//var respon  Respon;
	textBytes := []byte(body)
	respon := Respon{}
	jsonErr := json.Unmarshal(textBytes, &respon)
	if jsonErr != nil {
		//	log.Fatal(jsonErr)
	}

	return string(body)
}
