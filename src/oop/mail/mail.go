// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package mail contains the Mail API interfaces and datatypes for sending Emails.
package mail

type Message struct {
	To      string
	Subject string
	Text    string
}

// Sender is an interface to send mails.
type Sender interface {

	// Send a mail to a given address with a subject and text.
	Send(message Message) error
}

// END OMIT
