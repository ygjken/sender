package webhook

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/ygjken/webhook/pkg/reader"
)

type payload struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

func WebHook(hookurl string) error {
	// get username
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	// get stdin
	text := string(reader.ReadStdin())

	msg, err := json.Marshal(payload{Text: text, Username: name})
	if err != nil {
		return err
	}

	resp, err := http.PostForm(hookurl, url.Values{"payload": {string(msg)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
