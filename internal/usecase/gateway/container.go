package gateway

type Container struct {
	Authenticator    Authenticator
	Mailer           Mailer
	PluginRepository PluginRepository
	DataSource       DataSource
	Github           PluginRegistry
	File             File
}
