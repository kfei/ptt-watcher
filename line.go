package main

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func lineMessageGenerator(message NotificationMessage) string {
	body := fmt.Sprintf("New entries in *%s*\n", message.Subscription.Name)
	for _, item := range message.Items {
		body += fmt.Sprintf("\n%s\n%s\n", item.Title, item.Href)
	}
	return body
}

func lineNotifier() {
	channelSecret := config.Notifications.Line.ChannelSecret
	channelAccessToken := config.Notifications.Line.ChannelAccessToken
	toUserId := config.Notifications.Line.ToUserId

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatal("Unable to initialize Line bot")
	}

	for {
		select {
		case notification := <-nLineChan:
			messageBody := lineMessageGenerator(notification)
			subName := bold(notification.Subscription.Name)
			_message := linebot.NewTextMessage(messageBody)
			_, err := bot.PushMessage(toUserId, _message).Do()
			if err != nil {
				log.Fatalf("%s Error while sending Line message", subName, messageBody)
				// TODO: Add a message queue for retry sending
				continue
			}
			log.Printf("%s Message successfully sent to Line", subName)
		}
	}
}
