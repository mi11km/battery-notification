package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Message struct {
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	//IconUrl   string `json:"icon_url"`
}

func NewMessage(username, text, iconEmoji string) *Message {
	return &Message{username, text, iconEmoji}
}

func (message *Message) Notify(webhookURL string) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("action=json marshal, err=%v\n", err)
		return err
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(jsonMessage)}})
	if err != nil {
		log.Printf("action=post message to slack, err=%v\n", err)
		return err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("action=send message to slack, err=failed to send message, status=%d, contents=%v\n",
			resp.StatusCode, contents)
		return fmt.Errorf("failed to send message, status: %d", resp.StatusCode)
	}
	if err != nil {
		log.Printf("action=read response body, err=%v\n", err)
		return err
	}
	return nil
}
