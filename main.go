package main

import (
	"log"
	"os"

	"./newchannel"
)

func main() {
	service := newchannel.NewChannelNotificationService(os.Getenv("SLACK_TOKEN"), os.Getenv("SLACK_NEW_CHANNEL_NOTIFICATION_CHANNEL_ID"))
	e := service.Run()
	if e != nil {
		log.Fatalf("%+v", e)
	}
}
