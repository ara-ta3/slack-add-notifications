package newchannel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/websocket"
)

var rtmStartURL = "https://slack.com/api/rtm.start"

var slackAPIEndpoint = "https://slack.com/api/"

var origin = "http://localhost"

type rtmStartResponse struct {
	OK    bool   `json:"ok"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

type slackMessage struct {
	Type    string  `json:"type"`
	Channel channel `json:"channel"`
}

type channel struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Created interface{} `json:"created"`
	Creator string      `json:"creator"`
}

type slackClient struct {
	Token string
}

func (cli *slackClient) connectToRTM() (*websocket.Conn, error) {
	v := url.Values{
		"token": {cli.Token},
	}
	resp, e := http.Get(rtmStartURL + "?" + v.Encode())
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	byteArray, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	res := rtmStartResponse{}
	e = json.Unmarshal(byteArray, &res)
	if e != nil {
		return nil, e
	}
	if !res.OK {
		return nil, fmt.Errorf(res.Error)
	}
	ws, e := websocket.Dial(res.URL, "", origin)
	if e != nil {
		return nil, e
	}
	return ws, nil
}

func (cli *slackClient) polling(messageChan chan *slackMessage, errorChan chan error) {
	ws, e := cli.connectToRTM()
	if e != nil {
		errorChan <- e
		return
	}
	defer ws.Close()
	for {
		var msg = make([]byte, 1024)
		n, e := ws.Read(msg)
		if e != nil {
			errorChan <- e
		} else {
			message := slackMessage{}
			err := json.Unmarshal(msg[:n], &message)
			if err == nil {
				messageChan <- &message
			}
		}
	}
}

func (cli *slackClient) postMessage(channelID, text, userName, iconEmoji string) ([]byte, error) {
	res, e := http.PostForm(slackAPIEndpoint+"chat.postMessage", url.Values{
		"token":      {cli.Token},
		"channel":    {channelID},
		"text":       {text},
		"username":   {userName},
		"as_user":    {"false"},
		"icon_emoji": {iconEmoji},
		"link_names": {"0"},
	})
	if e != nil {
		return nil, e
	}
	defer res.Body.Close()
	byteArray, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}
	return byteArray, nil
}
