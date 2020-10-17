// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package client

import (
	"testing"

	"github.com/0xqab/concepts-of-programming-languages/oop/mail/smtp"
)

// configure Registry to know which mail implementation is used.
func init() {
	Registry.Register("mail.Sender", new(smtp.MailSenderImpl))
}

func TestMail(t *testing.T) {
	err := SendMail("johannes.weigend@qaware.de", "Hello", "Hello from Go!")
	if err != nil {
		t.Error("err should have been nil")
	}
}
