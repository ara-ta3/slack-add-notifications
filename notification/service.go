package notification

import (
	"fmt"

	"github.com/ara-ta3/slack-add-notifications/slack"
)

const (
	channelCreatedEventType = "channel_created"
	emojiChangedEventType   = "emoji_changed"
	emojiAddedSubEventType  = "add"
)

type NotificationService struct {
	SlackClient              slack.Client
	NewChannelNotificationID string
	NewEmojiNotificationID   string
	format                   PostMessageFormat
}

type PostMessageFormat struct {
	UserName  string  `json:"userName"`
	IconEmoji string  `json:"emoji"`
	Message   Message `json:"message"`
}

type Message struct {
	NewChannel string `json:"channel"`
	NewEmoji   string `json:"emoji"`
}

func NewNotificationService(slackClient slack.Client, newChannelNotificationID, newEmojiNotificationID string, format PostMessageFormat) NotificationService {
	return NotificationService{
		SlackClient:              slackClient,
		NewChannelNotificationID: newChannelNotificationID,
		NewEmojiNotificationID:   newEmojiNotificationID,
		format:                   format,
	}
}

func (service *NotificationService) Run() error {
	messageChan := make(chan *slack.SlackMessage)
	errorChan := make(chan error)

	go service.SlackClient.Polling(messageChan, errorChan)
	for {
		select {
		case msg := <-messageChan:
			if service.isNewChannelNotification(msg) {
				e := service.postNewChannel(msg.Channel.ID, msg.Channel.Name)
				if e != nil {
					return e
				}
			}

			if service.isNewEmojiNotification(msg) {
				e := service.postNewEmoji(msg.Name)
				if e != nil {
					return e
				}
			}
		case e := <-errorChan:
			return e
		default:
			break
		}
	}
	return nil
}

func (service *NotificationService) postNewChannel(channelID, channelName string) error {
	text := service.format.Message.NewChannel + fmt.Sprintf(" <#%s|%s>", channelID, channelName)
	r, e := service.SlackClient.PostMessage(
		service.NewChannelNotificationID,
		text,
		service.format.UserName,
		service.format.IconEmoji,
	)
	if e != nil {
		return e
	}
	fmt.Printf("%+v\n", string(r))
	return nil
}

func (service *NotificationService) postNewEmoji(emojiName string) error {
	text := service.format.Message.NewEmoji + fmt.Sprintf(" :%s:", emojiName)
	r, e := service.SlackClient.PostMessage(
		service.NewEmojiNotificationID,
		text,
		service.format.UserName,
		service.format.IconEmoji,
	)
	if e != nil {
		return e
	}
	fmt.Printf("%+v\n", string(r))
	return nil
}

func (service *NotificationService) isNewChannelNotification(m *slack.SlackMessage) bool {
	return m.Type == channelCreatedEventType
}

func (service *NotificationService) isNewEmojiNotification(m *slack.SlackMessage) bool {
	return m.Type == emojiChangedEventType && m.Subtype == emojiAddedSubEventType
}
