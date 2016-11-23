package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

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
		case message := <-nLineChan:
			_message := linebot.NewTextMessage(message.Body)
			_, err := bot.PushMessage(toUserId, _message).Do()
			if err != nil {
				log.Fatal("Error while sending Line message\n", message.Body)
				continue
			}
			log.Println("Message successfully sent to Line")
		}
	}
}
