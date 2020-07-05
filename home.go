package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//var path string
func path() string {
	path, err := os.Getwd()
	if err != nil {
		//	log.Println(err)
	}
	return path
}
var Available = false
var Availables bool
var HashSt string
var Country string
var ClientID string
var Secret_Key string
var Aus string
var Access_Token string
var authorize_code string
var URLq string
var UrlAuthorize string
var JWT_header_payload string
var Send_receive string
var Approv_Hold_Reject string
var URL_Admin string
var Envairment string = GetString("ENV")
var PublicKey string
var SecretKey string
var Currency string
var Currency2 string
var Message string
var Action string
var Lang string

// Make token
func Token() string {
	return ReqGetAccess_Token(URLq, Encoded(Aus))
}

type ViewData struct {
	Lang string
	Actions string
	Country string
	Availables bool
	Available bool
	JW        string
	ENV       string
	PublicKey string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := url.Parse(r.URL.String())

	if u.Query()["country"] != nil {
		Country = u.Query()["country"][0]
		fmt.Println("r.URL -- >>>   ", u.Query()["country"][0])
	} else {fmt.Println("Incorrect URL - country not sent")

	Message = "Incorrect URL - country not sent"
		Availables = true
	}

	if u.Query()["lang"] != nil {
		Lang = u.Query()["lang"][0]
		fmt.Println("r.URL -- >>>   ", u.Query()["lang"][0])
	} else {fmt.Println("Incorrect URL - lang not sent")
		Message = "Incorrect URL - lang not sent"
		Availables = true

	}

	if u.Query()["action"] != nil {
		Action = u.Query()["action"][0]
		fmt.Println("r.URL -- >>>   ", u.Query()["action"][0])
	} else {fmt.Println("Incorrect URL - action not sent")
		Message = "Incorrect URL - action not sent"
		Availables = true
	}



	switch {
	case strings.Contains(Country, "ukraine"):
		 URLq = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/1/oauth2/token"
		 UrlAuthorize = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/1/customers/authorize"

		//admin
		 ClientID = "UA713575"
		 Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		 Aus = ClientID + ":" + Secret_Key
		// Merchant settings
		 PublicKey = "CFe8x4Z"
		 SecretKey = "GVVFf7OevQ_tPuoMXFvm5vgGFQpgjFmq"
		Currency = "USD"
		 Currency2 = "USD"


	case strings.Contains(Country, "kenya"):
		URLq = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/1/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/1/customers/authorize"
		//admin
		ClientID = "KE932150"
		Secret_Key = "yQ@cTMmCoKCWshyA6#E7v2V0"
		Aus = ClientID + ":" + Secret_Key
		// Merchant settings
		PublicKey = "x2SZaJp"
		SecretKey = "6xrw38YLfIc4e!J2dw5a1OeE8JZOrBCS"
		Currency = "KES"
		Currency2 = "KES"


	default:
	}












	Access_Token = Token()
	fmt.Println("Access_Token -- >>>   ",Access_Token)
	JWT()
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	fmt.Println(r.Header)
	fmt.Println(r.Body)





	posts := ViewData{
		Lang: Lang,
		Actions: Action,
		Country: Country,
		Available: Availables,
		JW:        HashSt,
		ENV:       Envairment,
		PublicKey: PublicKey,
	}

	tmpl.Execute(w, posts)
}

type IdNum struct {
	IdNum string `json:"idnumber"`
	List  string `json:"list"`
}

//Change status transaction
func IndexApproveSend(w http.ResponseWriter, r *http.Request) {
	fmt.Println("qwqwqwqqwqw ---->>>>>>       ", r)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(requestBody)

	//Parse JSON
	textBytes := []byte(requestBody)
	respo2 := IdNum{}
	jsonErr := json.Unmarshal(textBytes, &respo2)
	if err != nil {
		fmt.Println(jsonErr)
	}

	// Select data from ajax
	Id_Number_ := respo2.IdNum
	List_ := respo2.List

	switch portalUrl := List_; { // missing expression means "true"
	case portalUrl == "notselected":
		Send_receive = ""
		Approv_Hold_Reject = ""

	case portalUrl == "ApproveSendTransaction":
		Send_receive = "send"
		Approv_Hold_Reject = "approve"

	case portalUrl == "HoldSendTransaction":
		Send_receive = "send"
		Approv_Hold_Reject = "suspend"

	case portalUrl == "RejectSendTransaction":
		Send_receive = "send"
		Approv_Hold_Reject = "reject"

	case portalUrl == "ApproveReceiveTransaction":
		Send_receive = "receive"
		Approv_Hold_Reject = "approve"

	case portalUrl == "HoldReceiveTransaction":
		Send_receive = "receive"
		Approv_Hold_Reject = "suspend"

	case portalUrl == "RejectReceiveTransaction":
		Send_receive = "receive"
		Approv_Hold_Reject = "reject"

	default:
		fmt.Println("Good evening!")
	}

	// Check data
	if Id_Number_ == "" || List_ == "notselected" {
		w.Write([]byte("<p><span style=\"color: #ff0000;\"><strong>Not all fields are filled</strong></span></p>"))
	} else {

		URL_Admin = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/1/" + Send_receive + "_transactions/" + Approv_Hold_Reject



		fmt.Println("Id_Number_   ----  >>>>       ",Id_Number_)
		fmt.Println("URL_Admin   ----  >>>>       ",URL_Admin)
		fmt.Println("Access_Token   ----  >>>>       ",Access_Token)

		w.Write([]byte(ChangeStatus(Id_Number_, URL_Admin, Access_Token)))
	}

}
