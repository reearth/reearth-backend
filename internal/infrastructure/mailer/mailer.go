package mailer

import (
	"net/smtp"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mailer struct {
	c gateway.MailerConfig
}

func InitMailer(conf gateway.MailerConfig) gateway.Mailer {
	return &mailer{
		c: conf,
	}
}

func (m *mailer) SendMail(to []string, content string) error {
	auth := smtp.PlainAuth("", m.c.Username, m.c.Password, m.c.Host)
	err := smtp.SendMail(m.c.Host+":"+m.c.Port, auth, m.c.Username, to, []byte(content))
	if err != nil {
		return err
	}
	return nil
}

func (m *mailer) SendGrid(toEmail, toName, subject, plainContent, htmlContent string) (*rest.Response, error) {
	sender := mail.NewEmail(m.c.Username, m.c.Email)
	receiver := mail.NewEmail(toName, toEmail)
	message := mail.NewSingleEmail(sender, subject, receiver, plainContent, htmlContent)
	client := sendgrid.NewSendClient(m.c.SendGridAPI)
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}
	return response, nil
}
