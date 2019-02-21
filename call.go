package tester

import (
	"encoding/json"

	"github.com/yields/phony/pkg/phony"
)

type Call interface {
	JSON() string
}

type Message struct {
	MessageID   string `json:"messageId"`
	Type        string `json:"type"`
	UserID      string `json:"userId"`
	AnonymousID string `json:"anonymousId"`
}

func NewMessage() *Message {
	m := Message{
		MessageID: phony.Get("id"),
	}
	return &m
}

type Track struct {
	Message
	Event      string                 `json:"event"`
	Properties map[string]interface{} `json:"properties"`
	Context    map[string]interface{} `json:"context"`
}

func NewTrack() *Track {
	t := Track{}
	t.Message = *NewMessage()
	t.Type = "track"
	return &t
}

func (t *Track) JSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
