package rice

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mailer interface {
	SendMail(subject, body string, to []string) error
}

type Mail struct {
	Auth smtp.Auth
	Host string
	Addr string
	From string
}

func NewMail(auth smtp.Auth, username, password, host, port string) *Mail {
	return &Mail{
		Auth: auth,
		Addr: fmt.Sprintf("%s:%s", host, port),
		From: username,
	}
}

func (m *Mail) SendMail(subject, body string, to []string) error {
	tos := strings.Join(to, ", ")
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", tos, subject, body)

	return smtp.SendMail(m.Addr, m.Auth, m.From, to, StringByteUnsafe(msg))
}
