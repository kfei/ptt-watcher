package main

import (
	"fmt"
	"log"

	"github.com/nlopes/slack"
)

func slackMessageGenerator(message NotificationMessage) string {
	body := fmt.Sprintf("New entries in *%s*\n", message.Subscription.Name)
	for _, item := range message.Items {
		body += fmt.Sprintf("><%s|%s>\n", item.Href, item.Title)
	}
	return body
}

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
		case notification := <-nSlackChan:
			messageBody := slackMessageGenerator(notification)
			_, _, err := api.PostMessage(channel, messageBody, params)
			if err != nil {
				log.Println("Error while sending Slack message:\n", messageBody)
				log.Println("Detail:\n", err)
				// TODO: Add a message queue for retry sending
				continue
			}
			log.Println("Message successfully sent to Slack")
		}
	}
}
