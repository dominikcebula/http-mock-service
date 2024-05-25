package main

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

type RuleHandler struct {
	rule Rule
}

func NewRuleHandlers(config Config) (ruleHandlers []RuleHandler) {
	for _, rule := range config.Rules {
		ruleHandlers = append(ruleHandlers, RuleHandler{rule})
	}
	return ruleHandlers
}

func (s *RuleHandler) CanHandle(request *http.Request) bool {
	requestPath := request.URL.Path
	matched, err := regexp.MatchString(s.rule.Request.Path, requestPath)
	if err != nil {
		log.Fatalf("Error occured during matching request path %s: %v", requestPath, err)
		return false
	}
	return matched
}

func (s *RuleHandler) Handle(response http.ResponseWriter, request *http.Request) {
	handleDelay(s)
	handleResponseCode(response, s)
	err := handleResponseBody(response, s)
	if err != nil {
		log.Fatalf("Error occured while writing response for request path %s: %v", request.URL.Path, err)
	}
}

func handleDelay(s *RuleHandler) {
	time.Sleep(s.rule.Response.Delay)
}

func handleResponseCode(response http.ResponseWriter, s *RuleHandler) {
	if s.rule.Response.Code > 0 {
		response.WriteHeader(s.rule.Response.Code)
	}
}

func handleResponseBody(response http.ResponseWriter, s *RuleHandler) error {
	_, err := response.Write([]byte(s.rule.Response.Body))
	return err
}
