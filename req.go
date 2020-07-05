package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//type Tokens struct {
//	Country,Domain, Path, Authorization string
//}

type respo struct {
	Access_Token string `json:"Access-Token"`
}

func ReqGetAccess_Token(q, w string) string {
	fmt.Println("func ReqGetAccess_Token")

	// Get Request
	req, err := http.NewRequest("GET", q, nil)
	if err != nil {
		fmt.Println("Error reading requests. ", err)
	}
	//
	req.Header.Set("Authorization", "Basic "+ w)
	client := &http.Client{Timeout: time.Second * 10000}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error reading response. ", err)
	}
	fmt.Println("Request ->   ", req)

	defer resp.Body.Close()

	// Получаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body. ", err)
	}
	fmt.Println("Response ->   ", resp.Body)

	//Parse JSON
	textBytes := []byte(body)
	respo1 := respo{}
	jsonErr := json.Unmarshal(textBytes, &respo1)
	if err != nil {
		fmt.Println(jsonErr)
	}
	//fmt.Println("Access_Token         ----------->>>>>>>>         " + respo1.Access_Token)

	//log.Infoln(resp.Header)
	//log.Infoln(resp.Body)

	return respo1.Access_Token
}
