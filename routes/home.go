package routes

import (
	log "github.com/Sirupsen/logrus"
	"html/template"
	"net/http"
)
type ViewData struct{
	Available bool
	JW string
	PublicKey string
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
