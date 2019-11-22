package routes

import (
	"fmt"
	"net/http"
	"html/template"
)
type ViewData struct{
	Available bool
}




func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("template/index.html"))

	fmt.Println("Alex \n")
	posts := ViewData{
		Available: true,
	}

	tmpl.Execute(w, posts)
}
