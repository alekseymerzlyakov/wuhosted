package routes

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/magiconair/properties"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

//var path string
func path () string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

//Open Properties file
func GetString(key string) string {
	p := properties.MustLoadFile(path() + "/wu.properties", properties.UTF8)
	return p.MustGet(key)
}

	var HashSt string
	var ClientID string = GetString("ClientID")
	var Secret_Key string = GetString("Secret_Key")
	var Aus string = ClientID + ":" + Secret_Key
	var Access_Token string
	var authorize_code string
	var URLq string = "https://" + GetString("URL") + "/admin/api/1/oauth2/token"
	var UrlAuthorize string = "https://" + GetString("URL") + "/admin/api/1/customers/authorize"
	var JWT_header_payload string
	var Send_receive string
	var Approv_Hold_Reject string
	var URL_Admin string


// Make token
func Token() string {
	return ReqGetAccess_Token(URLq, Encoded(Aus))
}

type ViewData struct{
	Available bool
	JW string
	PublicKey string
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	Access_Token = Token()
	JWT()
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	log.Info(r.Header)
	log.Info(r.Body)

	posts := ViewData{
		Available: false,
		JW: HashSt,
		PublicKey: GetString("publicKey"),
	}

	tmpl.Execute(w, posts)
}





type IdNum struct{
	IdNum string `json:"idnumber"`
	List string `json:"list"`
}
//Change status transaction
func IndexApproveSend(w http.ResponseWriter, r *http.Request)  {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info(err)
	}
	log.Debug(requestBody)

	//Parse JSON
	textBytes := []byte(requestBody)
	respo2 := IdNum{}
	jsonErr := json.Unmarshal(textBytes, &respo2)
	if err != nil {
		log.Println(jsonErr)
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

		URL_Admin = "https://" + GetString("URL") + "/admin/api/1/" + Send_receive + "_transactions/" + Approv_Hold_Reject

		w.Write([]byte(ChangeStatus(Id_Number_, URL_Admin, Access_Token)))
	}



}

