// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package client contains sample code for the mail components.
package client

import (
	"github.com/s-macke/concepts-of-programming-languages/src/oop/mail"
	"github.com/s-macke/concepts-of-programming-languages/src/oop/mail/util"
)

// Registry is the central configuration for the service locator
var Registry = util.NewRegistry()

// SendMail sends a mail to a receiver.
func SendMail(to, subject, text string) error {

	// Create an implementation for the mail.Sender interface.
	var sender = Registry.Get("mail.Sender").(mail.Sender)

	email := mail.Message{To: to, Subject: subject, Text: text}
	return sender.Send(email)
}

// EOF OMIT
