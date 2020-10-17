// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package mail contains the Mail API interfaces and datatypes for sending Emails.
package mail

type Email struct {
	From    string
	To      string
	Cc      string
	Bcc     string
	Subject string
	Message string
}

// Sender is a interface to send mails.
type Sender interface {

	// Send an email to a given address with a subject and message.
	Send(email Email) error
}

// END OMIT
