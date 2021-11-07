package mailer

import (
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"reflect"
	"testing"
)

func TestNewWithSMTP(t *testing.T) {
	type args struct {
		host     string
		port     string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want gateway.Mailer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithSMTP(tt.args.host, tt.args.port, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithSMTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_encodeContent(t *testing.T) {
	type fields struct {
		to           []string
		subject      string
		plainContent string
		htmlContent  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				to:           tt.fields.to,
				subject:      tt.fields.subject,
				plainContent: tt.fields.plainContent,
				htmlContent:  tt.fields.htmlContent,
			}
			got, err := m.encodeContent()
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encodeContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_encodeMessage(t *testing.T) {
	type fields struct {
		to           []string
		subject      string
		plainContent string
		htmlContent  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				to:           tt.fields.to,
				subject:      tt.fields.subject,
				plainContent: tt.fields.plainContent,
				htmlContent:  tt.fields.htmlContent,
			}
			got, err := m.encodeMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smtpMailer_SendMail(t *testing.T) {
	type fields struct {
		host     string
		port     string
		username string
		password string
	}
	type args struct {
		to           []gateway.Contact
		subject      string
		plainContent string
		htmlContent  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &smtpMailer{
				host:     tt.fields.host,
				port:     tt.fields.port,
				username: tt.fields.username,
				password: tt.fields.password,
			}
			if err := m.SendMail(tt.args.to, tt.args.subject, tt.args.plainContent, tt.args.htmlContent); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
