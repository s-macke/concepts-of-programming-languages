package v2

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type SenderImpl struct {
}

func (s SenderImpl) Send(message Message, timeout time.Duration) error {
	taskChan := make(chan bool)
	go func() {
		log.WithFields(log.Fields{
			"from":    message.From,
			"to":      message.To,
			"size":    len(message.Message),
			"timeout": timeout,
		}).Info("sending message with timeout...")
		time.Sleep(1000 * time.Millisecond)
		taskChan <- true
	}()
	timerChan := time.After(timeout)
	go func() {
		select {
		case <-taskChan:
			log.Info("message sent")
		case <-timerChan:
			log.Warn("timeout reached")
		}
	}()
	return nil
}
