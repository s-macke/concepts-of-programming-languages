package main

import (
	mail "github.com/0xqab/concepts-of-programming-languages/modules/mail/v2"
	log "github.com/sirupsen/logrus"
	"plugin"
	"time"
)

func main() {
	log.Info("opening plugin")
	p, err := plugin.Open("../plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	log.Info("looking up symbol")
	senderSymbol, err := p.Lookup("MailSender")
	if err != nil {
		panic(err)
	}
	sender := senderSymbol.(*mail.SenderImpl)

	log.Info("sending mail")
	err = sender.Send(mail.Message{
		From:    "bernhard.saumweber@qaware.de",
		To:      "johannes.weigend@qaware.de",
		Subject: "KP",
		Message: "Hallo Welt!",
	}, 2000*time.Millisecond)
	if err != nil {
		panic(err)
	}
	time.Sleep(2100 * time.Millisecond)
}
