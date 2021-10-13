package gateway

type Mailer interface {
	Send(subject string, body string, to ...string) error
}
