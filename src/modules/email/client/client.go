package main

import (
	v1 "github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.Info("sending mail")
	sender := v1.NewSender()
	err := sender.Send(v1.Message{
		From:    "sebastian.macke@qaware.de",
		To:      "johannes.weigend@qaware.de",
		Subject: "KP",
		Message: "Hallo Welt!",
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(2100 * time.Millisecond)
}
