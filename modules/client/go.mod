module github.com/0xqab/concepts-of-programming-languages/modules/client

go 1.15

require (
	github.com/0xqab/concepts-of-programming-languages/modules/mail/v2 v2.0.0
	github.com/sirupsen/logrus v1.7.0
)

replace github.com/0xqab/concepts-of-programming-languages/modules/mail/v2 => ../mail/v2
