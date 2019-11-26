package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// Get Request
	req, err := http.NewRequest("GET", q, nil)
	if err != nil {
		log.Fatal("Error reading requests. ", err)
	}
	//
	req.Header.Set("Authorization", "Basic " + w)
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()


	// Получаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}


	//Parse JSON
	textBytes := []byte(body)
	respo1 := respo{}
	jsonErr := json.Unmarshal(textBytes, &respo1)
	if err != nil {
		fmt.Println(jsonErr)
	}
	//fmt.Println("Access_Token         ----------->>>>>>>>         " + respo1.Access_Token)

	fmt.Printf("%s\n", body)
	return respo1.Access_Token
}

