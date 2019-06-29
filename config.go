package main

import (
	"encoding/json"
	"os"

	"github.com/ara-ta3/slack-add-notifications/notification"
)

type Config struct {
	SlackAPIToken string                         `json:"slackApiToken"`
	NotificateTo  NotificateTo                   `json:"notificationChannel"`
	Format        notification.PostMessageFormat `json:"format"`
	Debug         bool                           `json:"debug"`
}

type NotificateTo struct {
	NotificationChannelID string `json:"newChannel"`
	NotificationEmojiID   string `json:"newEmoji"`
}

func ReadConfig(path string) (*Config, error) {
	result := Config{}
	file, openErr := os.Open(path)
	if openErr != nil {
		return nil, openErr
	}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
