package routes

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
)
type ViewData struct{
	Available bool
	JW string
	PublicKey string
}

type IdNum struct{
	IdNum string `json:"idnumber"`
}


func IndexHandler(w http.ResponseWriter, r *http.Request)  {
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
	//fmt.Println("Access_Token         ----------->>>>>>>>         " + respo1.Access_Token)

	//log.Infoln(resp.Header)
	log.Infoln("\n QQQQQQQQQQQQQQQQQQQQQQQ                "   +  respo2.IdNum + "        ALEX")


	//w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(respo2.IdNum))

	//w.Write([]byte(textBytes))

}

