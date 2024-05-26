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
		s.noHandlerFoundResponse(response, request)
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

func (s *HttpService) noHandlerFoundResponse(response http.ResponseWriter, request *http.Request) {
	log.Printf("No rule handler found for request \"%s\"...", request.URL.Path)
	response.WriteHeader(501)
	_, err := response.Write([]byte("No handler found for request."))
	if err != nil {
		log.Panicf("Error writing response: %v", err)
	}
}
