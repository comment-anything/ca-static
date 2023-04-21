package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	dir := flag.String("static", "public", "Specify the directory to serve.")

	flag.Parse()

	fmt.Printf("Serving %s", *dir)

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	fmt.Printf("\nStarting static server on port 80.")

	log.Fatal(http.ListenAndServe(":80",
		http.DefaultServeMux))
}
