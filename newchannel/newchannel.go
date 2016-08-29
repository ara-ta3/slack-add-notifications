package newchannel

type ChannelNotificationService struct {
	SlackClient              slackClient
	NewChannelNotificationID string
}

func NewChannelNotificationService(slackAPIToken, newChannelNotificationID string) ChannelNotificationService {
	return ChannelNotificationService{
		SlackClient:              slackClient{Token: slackAPIToken},
		NewChannelNotificationID: newChannelNotificationID,
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
		case e := <-errorChan:
			return e
		default:
			break
		}
	}
	return nil
}

func (service *ChannelNotificationService) isTargetMessage(m *slackMessage) bool {
	return false
}
