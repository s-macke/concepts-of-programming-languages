// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package smtp sends mails over the smtp protocol.
package smtp

import (
	"github.com/0xqab/concepts-of-programming-languages/oop/mail"
	"log"
)

// MailSenderImpl is a sender object.
type MailSenderImpl struct {
}

// SendMail sends a mail to a receiver.
func (m *MailSenderImpl) Send(message mail.Message) error {
	log.Printf("Sending message with SMTP:\n  To: %v\n  Subject: %v\n  Text: %v\n",
		message.To, message.Subject, message.Text)
	return nil
}

//END OMIT
