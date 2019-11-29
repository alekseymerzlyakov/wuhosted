package routes

import (
	"fmt"
	"html/template"
	"net/http"
	log "github.com/Sirupsen/logrus"
)
type ViewData struct{
	Available bool
	JW string
}




func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HashSt   ------->>>>   QQQQQQQQ      >>>>>         ")
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	log.Info(r.Header)
	log.Info(r.Body)

	posts := ViewData{
		Available: false,
		JW: HashSt,
	}
	fmt.Println("HashSt   ------->>>>   QQQQQQQQ      >>>>>         " + posts.JW)

	tmpl.Execute(w, posts)
}
