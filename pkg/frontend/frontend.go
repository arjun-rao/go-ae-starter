package frontend

import (
	"html/template"
	"net/http"

	"github.com/arjun-rao/go-ae-starter/pkg/templates"
)

// HomeHandler renders the default page for the route '/'
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var pageTemplate = template.Must(template.ParseFiles(templates.Path("index.html")))
	data := struct {
		Title string
		Body  string
	}{
		Title: "Home",
		Body:  "Welocome to App Engine",
	}
	if err := pageTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
