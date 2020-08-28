package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("static/*.go.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home", nil)
}
func blogHandler(w http.ResponseWriter, r *http.Request) {

}
func reviewHandler(w http.ResponseWriter, r *http.Request) {

}
func methodHandler(w http.ResponseWriter, r *http.Request) {

}

//Middleware wrapper for mux.NewRouter()
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

//StartServer  starts server
func StartServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/blog", blogHandler)
	router.HandleFunc("/review", reviewHandler)
	router.HandleFunc("/method", methodHandler)
	router.PathPrefix("/stylesheets/").Handler(http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	return router
}
