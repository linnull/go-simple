package main

import (
	"log"
	"net/http"
)

func main() {

	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/root"))))

	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
