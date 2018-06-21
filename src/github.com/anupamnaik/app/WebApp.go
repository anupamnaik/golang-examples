package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/anupamnaik/app/functions"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GO WEB APP. And your lucky number is %d", functions.RandInt(100))
}

// WebApp the web application
func WebApp() {
	fmt.Println("Starting web application...")

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
