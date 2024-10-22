package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Server is running on port: 8082")

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
