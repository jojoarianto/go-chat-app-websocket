package domain

import "time"

type Message struct {
	Sender         string `json:"sender,omitempty"`
	ContentMessage string `json:"content_message"`
	CreatedAt      time.Time `json:"created_at"`
}

// NewMessage constructor to set default value
func NewMessage(msg string) Message {
	return Message{
		Sender:         "unknown",
		ContentMessage: msg,
		CreatedAt:      time.Now(),
	}
}
