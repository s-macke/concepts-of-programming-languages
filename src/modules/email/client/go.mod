module github.com/s-macke/concepts-of-programming-languages/src/modules/email/client

go 1.17

require (
	//github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail/v2 v2.0.0
	github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail v1.0.0
	github.com/sirupsen/logrus v1.7.1
)

require (
	github.com/magefile/mage v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace github.com/s-macke/concepts-of-programming-languages/src/modules/email/mail => ../mail
