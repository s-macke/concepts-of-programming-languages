package v2

import (
	"time"
)

type Message struct {
	From    string
	To      string
	Subject string
	Message string
}

type Sender interface {
	Send(Message, time.Duration) error
}

func NewSender() *SenderImpl {
	return &SenderImpl{}
}
