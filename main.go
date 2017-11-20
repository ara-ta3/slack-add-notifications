package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ara-ta3/slack-add-notifications/notification"
	"github.com/ara-ta3/slack-add-notifications/slack"
)

func main() {
	filePath := flag.String("c", "config.json", "file path to config.json")
	flag.Parse()
	fmt.Printf("config filepath: %s\n", *filePath)
	config, e := ReadConfig(*filePath)
	if e != nil {
		log.Fatalf("%+v", e)
	}

	service := notification.NewNotificationService(
		slack.Client{
			Token: config.SlackAPIToken,
		},
		config.NotificateTo.NotificationChannelID,
		config.NotificateTo.NotificationEmojiID,
		config.Format,
	)

	e = service.Run()
	if e != nil {
		log.Fatalf("%+v", e)
	}
}
