package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Coffee Blog: Service Booted up...\n")
	r := StartServer()
	http.Handle("/", Middleware(r))
	_, err := os.Stat(filepath.Join(".", "stylesheets", "sidebar.css"))
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
