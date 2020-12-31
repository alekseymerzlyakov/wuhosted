package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
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
var Device string
var HashSt string
var Country string
var WuHosted_Portal string
var ClientID string
var WuHosted string
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
var Envairment string
var PublicKey string
var SecretKey string
var Currency string
var Currency2 string
var Message string
var Action string
var Lang string
var CountryTeg string
var Env string

// Make token
func Token() string {
	return ReqGetAccess_Token(URLq, Encoded(Aus))
}

type ViewData struct {
	Lang       string
	WuHosted   string
	Actions    string
	Country    string
	Availables bool
	Device     string
	Available  bool
	JW         string
	ENV        string
	PublicKey  string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	u, _ := url.Parse(r.URL.String())

	if u.Query()["env"] != nil {
		Env = u.Query()["env"][0]
		Envairment = Env
		fmt.Println("\nenv -- >>>   ", u.Query()["env"][0])
	} else {
		fmt.Println("\nIncorrect env - this parameter mast contain pi or uat value")

		Message = "Incorrect env - this parameter mast contain pi or uat value"
		Availables = true
	}

	if u.Query()["country"] != nil {
		CountryTeg = u.Query()["country"][0]
		fmt.Println("\ncountry -- >>>   ", u.Query()["country"][0])
	} else {
		fmt.Println("I\nncorrect country - country not sent")

		Message = "\nIncorrect country - country not sent"
		Availables = true
	}

	if u.Query()["lang"] != nil {
		Lang = u.Query()["lang"][0]
		fmt.Println("\nlang -- >>>   ", u.Query()["lang"][0])
	} else {
		fmt.Println("\nIncorrect lang - lang not sent")
		Message = "\nIncorrect lang - lang not sent"
		Availables = true

	}

	if u.Query()["action"] != nil {
		Action = u.Query()["action"][0]
		fmt.Println("\naction -- >>>   ", u.Query()["action"][0])
	} else {
		fmt.Println("\nIncorrect action - action not sent")
		Message = "\nIncorrect action - action not sent"
		Availables = true
	}

	if u.Query()["currency"] != nil {
		Currency = u.Query()["currency"][0]
		fmt.Println("\ncurrency -- >>>   ", u.Query()["currency"][0])
	} else {
		fmt.Println("\nIncorrect currency - currency not sent")
		Message = "\nIncorrect currency - currency not sent"
		Availables = true
	}

	switch {

	case strings.EqualFold(CountryTeg, "morocco") && strings.EqualFold(Env, "pi"):
		WuHosted = "https://morocco." + Env + ".wuamerigo.com"
		WuHosted_Portal = "https://portal.morocco." + Env + ".wuamerigo.com"
		fmt.Println("\nCountry -- >>>   morocco")
		Country = "morocco"
		Device = "ios"

		URLq = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/customers/authorize"

		//admin
		//ClientID = "UA713575"
		//Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		//Aus = ClientID + ":" + Secret_Key
		// NEW Merchant settings
		PublicKey = "AvwMp"
		SecretKey = "pBuyj!84xD!chF9dUEKL4LXzjBsA8S"
		//Currency = "MAD"
		Currency2 = "MAD"

	case strings.EqualFold(CountryTeg, "morocco") && strings.EqualFold(Env, "uat"):
		WuHosted = "https://morocco." + Env + ".wuamerigo.com"
		WuHosted_Portal = "https://portal.morocco." + Env + ".wuamerigo.com"
		fmt.Println("\nCountry -- >>>   morocco")
		Country = "morocco"
		Device = "ios"

		URLq = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/customers/authorize"

		//admin
		//ClientID = "UA713575"
		//Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		//Aus = ClientID + ":" + Secret_Key
		// NEW Merchant settings
		PublicKey = "AvwMp"
		SecretKey = "pBuyj!84xD!chF9dUEKL4LXzjBsA8S"
		//Currency = "MAD"
		Currency2 = "MAD"

	case strings.EqualFold(CountryTeg, "ukraine_prod"):
		Country = "ukraine"
		Device = "ios"
		WuHosted = "https://ukraine-hosted.wu.com"
		WuHosted_Portal = "https://portal.ukraine-hosted.wu.com"
		fmt.Println("\nCountry -- >>>   ukraine_prod")
		URLq = "https://portal.ukraine-hosted.wu.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal.ukraine-hosted.wu.com/admin/api/2/customers/authorize"

		PublicKey = "_#P6t"
		SecretKey = "Su5qj@BoFmlDg@qEnOz_#h4AjmGQ6X"
		//Currency = "UAH"
		Currency2 = "UAH"

	case strings.EqualFold(CountryTeg, "ukraine") && strings.EqualFold(Env, "uat"):
		WuHosted = "https://ukraine." + Env + ".wuamerigo.com"
		WuHosted_Portal = "https://portal.ukraine." + Env + ".wuamerigo.com"
		fmt.Println("\nCountry -- >>>   ukraine")
		Country = "ukraine"
		Device = "ios"

		URLq = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/customers/authorize"

		//admin
		//ClientID = "UA713575"
		//Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		//Aus = ClientID + ":" + Secret_Key
		// NEW Merchant settings
		PublicKey = "3VGim"
		SecretKey = "NRZpznGPwBtWbyd5PFtpw@dOKNgdqA"
		//Currency = "USD"
		Currency2 = "USD"

	case strings.EqualFold(CountryTeg, "ukraine") && strings.EqualFold(Env, "pi"):
		WuHosted = "https://ukraine." + Env + ".wuamerigo.com"
		WuHosted_Portal = "https://portal.ukraine." + Env + ".wuamerigo.com"
		fmt.Println("\nCountry -- >>>   ukraine")
		Country = "ukraine"
		Device = "ios"

		URLq = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + Env + ".wuamerigo.com/admin/api/2/customers/authorize"

		//admin
		//ClientID = "UA713575"
		//Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		//Aus = ClientID + ":" + Secret_Key
		// NEW Merchant settings
		PublicKey = "ITXKG"
		SecretKey = "h1B3JD6a!roSklKOqgKiocP8cs9QGg"
		//Currency = "USD"
		Currency2 = "USD"

	case strings.EqualFold(CountryTeg, "ukraine_old"):
		WuHosted = "https://ukraine.uat.wuamerigo.com"
		WuHosted_Portal = "https://portal.ukraine.uat.wuamerigo.com"
		Country = "ukraine"
		Device = "ios"
		URLq = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/2/customers/authorize"

		//admin
		ClientID = "UA713575"
		Secret_Key = "uzB@LI4z#rF6QMiqGrnxDpPg"
		Aus = ClientID + ":" + Secret_Key
		// Merchant settings
		PublicKey = "vRpwp"
		SecretKey = "SvgcnbWOd9dkYRuJ@g5HMkQ2qwq7ih"
		//Currency = "USD"
		Currency2 = "USD"

	case strings.EqualFold(CountryTeg, "kenya"):
		Country = "kenya"
		Device = "web"
		WuHosted = "https://kenya.uat.wuamerigo.com"
		WuHosted_Portal = "https://portal.kenya.uat.wuamerigo.com"
		URLq = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/2/oauth2/token"
		UrlAuthorize = "https://portal." + Country + "." + GetString("ENV") + ".wuamerigo.com/admin/api/2/customers/authorize"
		//admin
		//ClientID = "KE932150"
		//Secret_Key = "JWorVd7A#Qk@aK936D3ZOmn!"
		//Aus = ClientID + ":" + Secret_Key
		// Merchant settings
		PublicKey = "iW0HE"
		SecretKey = "BCfPeJcSuk!WBuRZbF#WmQgyil6dz8"
		//Currency = "KES"
		Currency2 = "KES"

	default:

	}

	//Access_Token = Token()
	//fmt.Println("Access_Token -- >>>   ",Access_Token)

	//Authorize

	//authorize_code = Authorize(UrlAuthorize, Access_Token)
	//fmt.Println("authorize_code ---- >>>   ",authorize_code)

	fmt.Println("JWT")
	JWT()
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	fmt.Println("\nr.Header ->   ", r.Header)
	fmt.Println("\nr.Body ->   ", r.Body)
	log.Info("\nr.Header ->   ", r.Header)
	log.Info("\nr.Body ->   ", r.Body)

	posts := ViewData{
		Device:    Device,
		Lang:      Lang,
		WuHosted:  WuHosted,
		Actions:   Action,
		Country:   Country,
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

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("\nfunc IndexApproveSend r.URL ---->>>>>>       ", r.URL)
	log.Println("\nfunc IndexApproveSend r.Header ---->>>>>>       ", r.Header)
	log.Println("\nfunc IndexApproveSend: Request Body ---->>>>>>       ", string(requestBody))
	fmt.Println("\nfunc IndexApproveSend r.URL ---->>>>>>       ", r.URL)
	fmt.Println("\nfunc IndexApproveSend r.Header ---->>>>>>       ", r.Header)
	fmt.Println("\nfunc IndexApproveSend: Request Body ---->>>>>>       ", string(requestBody))
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

		URL_Admin = "/api/2/" + Send_receive + "_transactions/" + Approv_Hold_Reject

		fmt.Println("\nId_Number_   ----  >>>>       ", Id_Number_)
		fmt.Println("\nURL_Admin   ----  >>>>       ", URL_Admin)
		fmt.Println("\nAccess_Token   ----  >>>>       ", Access_Token)
		log.Println("\nId_Number_   ----  >>>>       ", Id_Number_)
		log.Println("\nURL_Admin   ----  >>>>       ", URL_Admin)
		log.Println("\nAccess_Token   ----  >>>>       ", Access_Token)

		w.Write([]byte(ChangeStatus(Id_Number_, URL_Admin)))
	}

}
