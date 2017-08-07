package frontend

import (
	"net/http"

	"github.com/arjun-rao/go-ae-starter/pkg/frontend"

	"github.com/gorilla/mux"
)

// Function that defines the routes in the application
func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", frontend.HomeHandler)

	// Route all requests through the Mux
	http.Handle("/", r)

}
