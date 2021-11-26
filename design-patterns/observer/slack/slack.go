package slack

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	apiURL  = "https://hooks.slack.com/services"
	webhook = "/T207WP78B/B012XGPKUSZ/cKhZjr5JhWRmkp2bGX9qPQQ0"
)

type Slack struct{}

func (s *Slack) Notify(data string) { sendMessage(data) }

func sendMessage(data string) {
	msg := fmt.Sprintf(`{"text":%q}`, data)
	req, err := http.NewRequest(http.MethodPost,
		apiURL+webhook,
		bytes.NewReader([]byte(msg)))
	if err != nil {
		log.Fatalf("the request could not be created: %v", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("the messaged could not be sent to slack.")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("the status code was not 200: %s", resp.Status)
	}
}
