package mailer

import (
	"net/smtp"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
)

type smtpMailer struct {
	host     string
	port     string
	username string
	password string
}

func NewWithSMTP(host, port, username, password string) gateway.Mailer {
	return &smtpMailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (m *smtpMailer) SendMail(to []gateway.Contact, subject, plainContent, htmlContent string) error {
	emails := make([]string, 0, len(to))
	for _, c := range to {
		emails = append(emails, c.Email)
	}
	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	var content string
	if len(htmlContent) > 0 {
		content = htmlContent
	} else {
		content = plainContent
	}
	err := smtp.SendMail(m.host+":"+m.port, auth, m.username, emails, []byte(content))
	if err != nil {
		return err
	}
	return nil
}
