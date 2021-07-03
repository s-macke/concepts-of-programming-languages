package main

import (
	mail "github.com/s-macke/concepts-of-programming-languages/modules/mail/v2"
)

var MailSender mail.SenderImpl

func main() {
	MailSender = *mail.NewSender()
}
