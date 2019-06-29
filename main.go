package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("debug option: %t\n", config.Debug)

	messageChan := make(chan *slack.SlackMessage)
	errorChan := make(chan error)
	service := notification.NewNotificationService(
		slack.Client{
			Token: config.SlackAPIToken,
		},
		config.NotificateTo.NotificationChannelID,
		config.NotificateTo.NotificationEmojiID,
		config.Format,
		messageChan,
		errorChan,
	)

	mux := http.NewServeMux()
	handle := notification.NewHandler(messageChan)
	mux.Handle("/", handle)
	if config.Debug {
		go func(mux *http.ServeMux, errorChan chan error) {
			e := http.ListenAndServe(":8080", mux)
			errorChan <- e
		}(mux, errorChan)
	}
	e = service.Run()
	if e != nil {
		log.Fatalf("%+v", e)
	}
}
