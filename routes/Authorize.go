package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Respon struct {
	Authoriz    string `json:"authorization_code"`
	Action_info string `json:"action_info"`
}

// Post Request
func Authorize(url, Access_Token string) string {

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
  	"name": "` + GetString("name") + `",
  	"currency": "` + GetString("currency") + `",
  	"balance": "` + GetString("balance") + `"
	},
		{
  	"account_number": "` + GetString("account_number2") + `",
  	"name": "` + GetString("name2") + `",
  	"currency": "` + GetString("currency2") + `",
  	"balance": "` + GetString("balance2") + `"
	}
  ]
}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error reading requests. ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Access_Token)

	// Create and Add cookie to request
	//cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	//req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10000}

	// Validate cookie and headers are attached
	//fmt.Println(req.Cookies())
	log.Info(req.Header)
	log.Info(req.Body)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	log.Info("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	log.Info("%s\n", body)
	//Parse JSON

	//var respon  Respon;
	textBytes := []byte(body)
	respon := Respon{}
	jsonErr := json.Unmarshal(textBytes, &respon)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return respon.Authoriz
}
