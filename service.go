package main

import (
	"log"
	"net/http"
)

type HttpService struct {
}

func (s *HttpService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("Handling http request...")

	_, err := response.Write([]byte("test\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Request handled.")
}
