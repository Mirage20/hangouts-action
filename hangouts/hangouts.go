package hangouts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hangouts struct {
	*http.Client
	URL string
}

func NewWebhookClient(url string) *Hangouts {
	return &Hangouts{
		Client: &http.Client{},
		URL:    url,
	}
}

func (h *Hangouts) Send(threadKey string, msg *Message) (*Message, error) {
	url := h.URL
	if len(threadKey) > 0 {
		url = fmt.Sprintf("%s&threadKey=%s", url, threadKey)
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	resp, err := h.Client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("post message error: %s", body)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rMsg := &Message{}
	err = json.Unmarshal(body, rMsg)
	if err != nil {
		return nil, err
	}
	return rMsg, nil
}
