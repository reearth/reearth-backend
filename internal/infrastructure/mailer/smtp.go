package mailer

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"strings"

	"github.com/reearth/reearth-backend/internal/usecase/gateway"
)

type smtpMailer struct {
	host     string
	port     string
	username string
	password string
}

type message struct {
	to           []string
	subject      string
	plainContent string
	htmlContent  string
}

func (m *message) encodeContent() (string, error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	newBoundary := "RELATED-" + boundary
	relatedBuffer, err := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"multipart/related; boundary=" + newBoundary}})
	if err != nil {
		return "", err
	}
	relatedWriter := multipart.NewWriter(relatedBuffer)
	err = relatedWriter.SetBoundary(fix70Boundary(newBoundary))
	if err != nil {
		return "", err
	}
	newBoundary = "ALTERNATIVE-" + newBoundary

	altBuffer, err := relatedWriter.CreatePart(textproto.MIMEHeader{"Content-Type": {"multipart/alternative; boundary=" + newBoundary}})
	if err != nil {
		return "", err
	}
	altWriter := multipart.NewWriter(altBuffer)
	err = altWriter.SetBoundary(fix70Boundary(newBoundary))
	if err != nil {
		return "", err
	}
	var content io.Writer
	content, err = altWriter.CreatePart(textproto.MIMEHeader{"Content-Type": {"text/html"}})
	if err != nil {
		return "", err
	}

	_, err = content.Write([]byte(m.htmlContent + "\r\n\r\n"))
	if err != nil {
		return "", err
	}
	content, err = altWriter.CreatePart(textproto.MIMEHeader{"Content-Type": {"text/plain"}})
	if err != nil {
		return "", err
	}
	_, err = content.Write([]byte(m.plainContent + "\r\n"))
	if err != nil {
		return "", err
	}
	_ = altWriter.Close()
	_ = relatedWriter.Close()
	return buf.String(), nil
}

func (m *message) encodeMessage() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.to, ",")))
	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	content, err := m.encodeContent()
	if err != nil {
		return nil, err
	}
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n\n %s", boundary, content))

	return buf.Bytes(), nil
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

	msg := &message{
		to:           emails,
		subject:      subject,
		plainContent: plainContent,
		htmlContent:  htmlContent,
	}

	encodedMsg, err := msg.encodeMessage()
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", m.username, m.password, m.host)

	return smtp.SendMail(m.host+":"+m.port, auth, m.username, emails, encodedMsg)
}

func fix70Boundary(b string) string {
	if len(b) > 70 {
		return b[0:69]
	}
	return b
}
