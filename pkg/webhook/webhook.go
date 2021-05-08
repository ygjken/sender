package webhook

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type payload struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

func post(webhookurl string) error {
	msg, err := json.Marshal(payload{Text: "Done", Username: "macmini"})
	if err != nil {
		return err
	}

	resp, err := http.PostForm(webhookurl, url.Values{"payload": {string(msg)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func WebHook() {
}
