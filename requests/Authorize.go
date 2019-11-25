package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	"first_name": " Oleksii ",
	"last_name": " Merzliakov ",
	"street1": " street semafor4 ",
	"city": " city Dnepr ",
	"zip_code": " 12322 ",
	"country": "KE",
	"phone_country_code": "380",
	"phone_number": " 993879792 ",
	"birth": "1975-01-01",
	"country_of_birth": "KE",
	"id_type": "A",
    	"id_number": " 111211116 ",
	"id_issuer": "KE",
	"id_expires": true,
	"id_expiration": "2020-12-23"
  },
  "bank_accounts": [
	{
  	"account_number": "223456123456",
  	"name": "Nata",
  	"currency": "KES",
  	"balance": "3323434"
	},
		{
  	"account_number": "993456123456",
  	"name": "Nata2",
  	"currency": "KES",
  	"balance": "9"
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
	//req.Header.Set("Host", "portal.kenya.uat.wuamerigo.com")

	// Create and Add cookie to request
	//cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	//req.AddCookie(&cookie)

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10000}

	// Validate cookie and headers are attached
	//fmt.Println(req.Cookies())
	fmt.Println(req.Header)
	fmt.Println("\n")
	fmt.Println(req.Body)
	fmt.Println("\n")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("\n\n")
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("\n\n")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
	//Parse JSON


	var respo1  Respon;
	//textBytes := []byte(body)
	//respo1 := Respon{}
	jsonErr := json.Unmarshal(body, &respo1.Authoriz)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		panic(http.StatusNotAcceptable)
	}

	fmt.Printf("authorization_code     =======>>>>>>>>        " +  respo1.Authoriz)
	fmt.Printf("action_info            =======>>>>>>>>        " +  respo1.Action_info)
	return respo1.Authoriz
}