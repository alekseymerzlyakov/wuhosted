package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Respones struct {
	Authoriz    string `json:"authorization_code"`
	Action_info string `json:"action_info"`
}

// Post Request
func ChangeStatus(IdNum, UrlAuthorize, Access_Token string) string {

	data := []byte(`[{"ID":"` + IdNum + `"}]`)
	//log.Info("request -> " + string(data))
	req, err := http.NewRequest("PUT", UrlAuthorize, bytes.NewBuffer(data))
	if err != nil {
		//	log.Fatal("Error reading requests. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + Access_Token)

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
