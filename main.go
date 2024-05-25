package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Creating HTTP Mock Service...")

	httpService := HttpService{}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: &httpService,
	}

	log.Println("Listening for connections...")
	log.Fatal(httpServer.ListenAndServe())
}
