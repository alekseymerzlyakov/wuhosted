package requests

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/alekseymerzlyakov/wuhosted/routes"
	"io/ioutil"
	"net/http"
	"time"
)

type Respon struct {
	Authoriz string `json:"authorization_code"`
	Action_info string `json:"action_info"`
}



// Post Request
func Authorize(url, Access_Token string) string {



	data := []byte(`{
  "customer_id": "526321446",
  "profile": {
	"first_name": "` + routes.GetString("first_name") + `",
	"last_name": "` + routes.GetString("last_name") + `",
	"street1": "` + routes.GetString("street1") + `",
	"city": "` + routes.GetString("city") + `",
	"zip_code": "` + routes.GetString("zip_code") + `",
	"country": "` + routes.GetString("country") + `",
	"phone_country_code": "` + routes.GetString("phone_country_code") + `",
	"phone_number": "` + routes.GetString("phone_number") + `",
	"birth": "` + routes.GetString("birth") + `",
	"country_of_birth": "` + routes.GetString("country_of_birth") + `",
	"id_type": "` + routes.GetString("id_type") + `",
    	"id_number": "` + routes.GetString("id_number") + `",
	"id_issuer": "` + routes.GetString("id_issuer") + `",
	"id_expires": ` + routes.GetString("id_expires") + `,
	"id_expiration": "` + routes.GetString("id_expiration") + `"
  },
  "bank_accounts": [
	{
  	"account_number": "` + routes.GetString("account_number") + `",
  	"name": "` + routes.GetString("name") + `",
  	"currency": "` + routes.GetString("currency") + `",
  	"balance": "` + routes.GetString("balance") + `"
	},
		{
  	"account_number": "` + routes.GetString("account_number2") + `",
  	"name": "` + routes.GetString("name2") + `",
  	"currency": "` + routes.GetString("currency2") + `",
  	"balance": "` + routes.GetString("balance2") + `"
	}
  ]
}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error reading requests. ", err)
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