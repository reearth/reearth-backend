package gateway

import "github.com/sendgrid/rest"

type MailerConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Email       string
	SendGridAPI string
}

type Mailer interface {
	SendMail(to []string, content string) error
	SendGrid(toEmail, toName, subject, plainContent, htmlContent string) (*rest.Response, error)
}
