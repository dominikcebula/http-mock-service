package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"
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
	return s.CanHandleRequestPath(request) && s.CanHandleHttpMethod(request)
}

func (s *RuleHandler) CanHandleRequestPath(request *http.Request) bool {
	requestPath := request.URL.Path
	matched, err := regexp.MatchString(s.rule.Request.Path, requestPath)
	if err != nil {
		log.Panicf("Error occured during matching request path %s: %v\n", requestPath, err)
		return false
	}
	return matched
}

func (s *RuleHandler) CanHandleHttpMethod(request *http.Request) bool {
	if len(s.rule.Request.Method) == 0 {
		return true
	}

	return strings.EqualFold(strings.TrimSpace(request.Method), strings.TrimSpace(s.rule.Request.Method))
}

func (s *RuleHandler) Handle(response http.ResponseWriter, request *http.Request) {
	s.handleDelay()
	s.handleResponseHeaders(response)
	s.handleResponseCode(response)
	err := s.handleResponseBody(response)
	if err != nil {
		log.Panicf("Error occured while writing response for request path %s: %v", request.URL.Path, err)
	}
}

func (s *RuleHandler) handleDelay() {
	time.Sleep(s.rule.Response.Delay)
}

func (s *RuleHandler) handleResponseHeaders(response http.ResponseWriter) {
	for headerName, headerValue := range s.rule.Response.Headers {
		response.Header().Add(headerName, headerValue)
	}
}

func (s *RuleHandler) handleResponseCode(response http.ResponseWriter) {
	if s.rule.Response.Code > 0 {
		response.WriteHeader(s.rule.Response.Code)
	}
}

func (s *RuleHandler) handleResponseBody(response http.ResponseWriter) error {
	_, err := response.Write([]byte(s.rule.Response.Body))
	return err
}
