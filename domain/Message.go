package domain

import (
	"time"

	"github.com/rs/xid"
)

type Message struct {
	ID			   xid.ID `json:"id"`
	Sender         string `json:"sender,omitempty"`
	ContentMessage string `json:"content_message"`
	CreatedAt      time.Time `json:"created_at"`
}

// NewMessage constructor to set default value
func NewMessage(msg string) Message {
	return Message{
		ID: 			xid.New(), 
		Sender:         "Anonim",
		ContentMessage: msg,
		CreatedAt:      time.Now(),
	}
}
