package main

import (
	"github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail/v2"
	log "github.com/sirupsen/logrus"
	"plugin"
	"time"
)

func main() {
	log.Info("opening plugin")
	p, err := plugin.Open("plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	log.Info("looking up symbol")
	senderSymbol, err := p.Lookup("MailSender")
	if err != nil {
		panic(err)
	}
	sender := senderSymbol.(*v2.SenderImpl)

	log.Info("sending mail")
	err = sender.Send(v2.Message{
		From:    "sebastian.macke@qaware.de",
		To:      "johannes.weigend@qaware.de",
		Subject: "KP",
		Message: "Hallo Welt!",
	}, 2000*time.Millisecond)
	if err != nil {
		panic(err)
	}
	time.Sleep(2100 * time.Millisecond)
}
