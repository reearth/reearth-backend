package mailer

import (
	"github.com/kataras/go-mailer"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
)

type customMailer struct {
	mailer *mailer.Mailer
}

func NewMailer(conf mailer.Config) gateway.Mailer {
	return &customMailer{
		mailer: mailer.New(conf),
	}
}

func (m *customMailer) Send(subject string, body string, to ...string) error {
	return m.mailer.Send(subject, body, to...)
}
