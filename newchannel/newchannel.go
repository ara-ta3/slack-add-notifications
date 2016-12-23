package newchannel

import "fmt"

var channelCreatedEventType = "channel_created"

type ChannelNotificationService struct {
	SlackClient              slackClient
	NewChannelNotificationID string
	format                   PostMessageFormat
}

type PostMessageFormat struct {
	UserName  string `json:"userName"`
	Text      string `json:"text"`
	IconEmoji string `json:"emoji"`
}

func NewChannelNotificationService(slackAPIToken, newChannelNotificationID string, format PostMessageFormat) ChannelNotificationService {
	return ChannelNotificationService{
		SlackClient:              slackClient{Token: slackAPIToken},
		NewChannelNotificationID: newChannelNotificationID,
		format: format,
	}
}

func (service *ChannelNotificationService) Run() error {
	messageChan := make(chan *slackMessage)
	errorChan := make(chan error)

	go service.SlackClient.polling(messageChan, errorChan)
	for {
		select {
		case msg := <-messageChan:
			if !service.isTargetMessage(msg) {
				continue
			}
			text := service.format.Text + fmt.Sprintf(" <#%s|%s>", msg.Channel.ID, msg.Channel.Name)
			r, e := service.SlackClient.postMessage(
				service.NewChannelNotificationID,
				text,
				service.format.UserName,
				service.format.IconEmoji,
			)
			if e != nil {
				return e
			}
			fmt.Printf("%+v\n", string(r))
		case e := <-errorChan:
			return e
		default:
			break
		}
	}
	return nil
}

func (service *ChannelNotificationService) isTargetMessage(m *slackMessage) bool {
	return m.Type == channelCreatedEventType
}
