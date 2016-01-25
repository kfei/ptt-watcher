package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
)

var config Config
var nChan = make(chan Notification)

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error while opening config.json", err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatal("Error while parsing config.json", err)
	}

	// Start Slack notifier
	go notifier()

	// Start feed watchers
	for _, sub := range config.Subscriptions {
		watcher(sub)
	}

	// Stop when hitting Ctrl-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	for {
		select {
		case <-sigChan:
			log.Println("Interrupted")
			return
		}
	}
}
