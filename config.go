package main

import (
	"encoding/json"
	"os"

	"./newchannel"
)

type Config struct {
	SlackAPIToken         string                       `json:"slackApiToken"`
	NotificationChannelID string                       `json:"notificationChannelID"`
	Format                newchannel.PostMessageFormat `json:"format"`
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
