package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("static/*.gohtml"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page")
	templates.ExecuteTemplate(w, "home", nil)
}
func blogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Blog page")
	templates.ExecuteTemplate(w, "blog", nil)

}
func reviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Review page")
	templates.ExecuteTemplate(w, "review", nil)

}
func methodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Methods page")
	templates.ExecuteTemplate(w, "methods", nil)
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
	router.HandleFunc("/methods", methodHandler)
	router.PathPrefix("/stylesheets/").Handler(http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	return router
}
