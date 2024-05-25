package main

import (
	"log"
	"net/http"
)

func main() {
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

	httpServer := &http.Server{
		Addr:    config.Server.Address,
		Handler: &httpService,
	}
	log.Printf("Listening for connections on address \"%s\"...\n", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}
