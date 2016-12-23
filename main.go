package main

import (
	"flag"
	"fmt"
	"log"

	"./newchannel"
)

func main() {
	filePath := flag.String("c", "config.json", "file path to config.json")
	flag.Parse()
	fmt.Printf("config filepath: %s\n", *filePath)
	config, e := ReadConfig(*filePath)
	if e != nil {
		log.Fatalf("%+v", e)
	}

	service := newchannel.NewChannelNotificationService(
		config.SlackAPIToken,
		config.NotificationChannelID,
		config.Format,
	)

	e = service.Run()
	if e != nil {
		log.Fatalf("%+v", e)
	}
}
