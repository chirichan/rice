package rice

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	GmailSmtp    = "smtp.gmail.com"
	GmailTLSPort = 587
	GmailSSLPort = 465
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

func NewMail(auth smtp.Auth, from, host string, port int) Mailer {
	return &Mail{
		Auth: auth,
		Addr: fmt.Sprintf("%s:%d", host, port),
		From: from,
	}
}

func (m *Mail) SendMail(subject, body string, to []string) error {
	tos := strings.Join(to, ", ")
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", tos, subject, body)

	return smtp.SendMail(m.Addr, m.Auth, m.From, to, StringByteUnsafe(msg))
}

func NewGmailAuth(username, password string) smtp.Auth {
	return smtp.PlainAuth("", username, password, GmailSmtp)
}
