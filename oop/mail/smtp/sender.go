// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package smtp sends mails over the smtp protocol.
package smtp

import (
	"log"

	"github.com/0xqab/concepts-of-programming-languages/oop/mail"
)

// MailSenderImpl is a sender object.
type MailSenderImpl struct {
}

// SendMail sends a mail to a receiver.
func (m *MailSenderImpl) Send(email mail.Email) error {
	log.Println("Sending message with SMTP to " + email.To + " message: " + email.Message)
	return nil
}

//END OMIT
