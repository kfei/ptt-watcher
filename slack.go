package main

import (
	"log"

	"github.com/nlopes/slack"
)

func slackNotifier() {
	token := config.Notifications.Slack.Token
	channel := config.Notifications.Slack.Channel
	userName := config.Notifications.Slack.UserName

	api := slack.New(token)
	params := slack.PostMessageParameters{}
	if len(userName) >= 0 {
		params.Username = userName
	}
	params.Markdown = true

	for {
		select {
		case message := <-nSlackChan:
			_, _, err := api.PostMessage(channel, message.Body, params)
			if err != nil {
				log.Println("Error while sending Slack message:\n", message.Body)
				log.Println("Detail:\n", err)
				// TODO: Add a message queue for retry sending
				continue
			}
			log.Println("Message successfully sent to Slack")
		}
	}
}
