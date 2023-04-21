package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Printf("\nStarting static server on port 80.")

	log.Fatal(http.ListenAndServe(":80",
		http.DefaultServeMux))
}
