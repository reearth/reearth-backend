package mailer

import (
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridMailer struct {
	senderName string
	email      string
	api        string
}

func NewWithSendGrid(senderName, email, api string) gateway.Mailer {
	return &sendgridMailer{
		senderName: senderName,
		email:      email,
		api:        api,
	}
}

func (m *sendgridMailer) SendMail(to []gateway.Contact, subject, plainContent, htmlContent string) error {
	contact := to[0]
	sender := mail.NewEmail(m.senderName, m.email)
	receiver := mail.NewEmail(contact.Name, contact.Email)
	message := mail.NewSingleEmail(sender, subject, receiver, plainContent, htmlContent)
	client := sendgrid.NewSendClient(m.api)
	_, err := client.Send(message)
	return err
}
