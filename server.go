package main

import (
	"log"
	"net/http"
	"strconv"
)

var httpServer *http.Server

func createServer() {
	log.Println("Reading config file...")
	config, err := readConfig()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Config file loaded.")

	log.Println("Creating request handler(s)...")
	ruleHandlers := NewRuleHandlers(config)
	log.Printf("Created %d request handler(s).\n", len(ruleHandlers))

	log.Println("Creating HTTP Mock Server...")
	httpService := HttpService{ruleHandlers}

	httpServer = &http.Server{
		Addr:    config.Server.Host + ":" + strconv.Itoa(config.Server.Port),
		Handler: &httpService,
	}
}

func startServer() {
	if httpServer == nil {
		log.Panicln("HTTP Server needs to be created first before executing operations against the server")
	}

	log.Printf("Listening for connections on address \"%s\"...\n", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}
