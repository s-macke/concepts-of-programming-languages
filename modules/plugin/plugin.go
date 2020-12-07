package main

import (
	mail "github.com/0xqab/concepts-of-programming-languages/modules/mail/v2"
)

var MailSender mail.SenderImpl

func main() {
	MailSender = *mail.NewSender()
}
