package notification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ara-ta3/slack-add-notifications/slack"
)

type DebugHandler struct {
	messageChannel chan *slack.SlackMessage
}

func NewHandler(messageChannel chan *slack.SlackMessage) *DebugHandler {
	return &DebugHandler{
		messageChannel: messageChannel,
	}
}

func (d *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		return
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var m *slack.SlackMessage
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d.messageChannel <- m
	w.WriteHeader(http.StatusNoContent)
}
