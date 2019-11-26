package routes

import (
	"fmt"
	"html/template"
	"net/http"
)
type ViewData struct{
	Available bool
	JW string
}




func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("template/index.html"))


	posts := ViewData{
		Available: false,
		JW: HashSt,
	}
	fmt.Println("HashSt   ------->>>>   QQQQQQQQ      >>>>>         " + posts.JW)

	tmpl.Execute(w, posts)
}
