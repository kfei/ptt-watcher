package main

import (
	"log"

	"github.com/nlopes/slack"
)

func notifier() {
	api := slack.New(config.SlackToken)
	params := slack.PostMessageParameters{}
	if len(config.SlackUserName) >= 0 {
		params.Username = config.SlackUserName
	}
	params.Markdown = true

	for {
		select {
		case notification := <-nChan:
			_, _, err := api.PostMessage(config.SlackChannel, notification.Body, params)
			if err != nil {
				log.Println("Error while sending Slack message:\n", notification.Body)
				log.Println("Detail:\n", err)
				// TODO: Add a message queue for retry sending
				continue
			}
			log.Println("Message successfully sent")
		}
	}
}
