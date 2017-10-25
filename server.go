package main

import (
	"log"
	"net/http"
)

func main() {
	assets := http.StripPrefix("/", http.FileServer(http.Dir("assets/")))
	http.Handle("/", assets)

	log.Println("Listening on port 8080.")
	http.ListenAndServe(":8080", nil)

}
