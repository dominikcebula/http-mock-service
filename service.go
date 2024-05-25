package main

import (
	"log"
	"net/http"
)

type HttpService struct {
	ruleHandlers []RuleHandler
}

func (s *HttpService) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("Handling http request \"%s\"...", request.URL.Path)

	ruleHandler := s.findHandler(request)
	if ruleHandler != nil {
		ruleHandler.Handle(response, request)
	} else {
		log.Printf("No rule handler found for request \"%s\"...", request.URL.Path)
	}

	log.Println("HTTP Request handled.")
}

func (s *HttpService) findHandler(request *http.Request) *RuleHandler {
	for _, handler := range s.ruleHandlers {
		if handler.CanHandle(request) {
			return &handler
		}
	}

	return nil
}
