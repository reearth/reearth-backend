package gateway

type Container struct {
	Authenticator Authenticator
	Mailer        Mailer
	DataSource    DataSource
	File          File
	Google        Google
}
