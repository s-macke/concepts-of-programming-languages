module github.com/s-macke/concepts-of-programming-languages/src/modules/email/plugin

go 1.17

require github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail/v2 v2.0.0

require (
	github.com/sirupsen/logrus v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail/v2 => ./../mail/v2