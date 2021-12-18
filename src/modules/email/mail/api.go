package mail

type Message struct {
	From    string
	To      string
	Subject string
	Message string
}

type Sender interface {
	Send(Message) error
}

func NewSender() *SenderImpl {
	return &SenderImpl{}
}
