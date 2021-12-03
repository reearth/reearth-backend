package email

import (
	_ "embed"
	htmlTmpl "html/template"
	textTmpl "text/template"
)

type Contents map[ContentType]*Content
type ContentType string

const (
	VerifyUser    ContentType = "VerifyUser"
	ResetPassword ContentType = "ResetPassword"
)

var (
	//go:embed templates/password_reset_html.tmpl
	passwordResetHTMLTMPL string
	//go:embed templates/password_reset_text.tmpl
	passwordResetTextTMPL string
	//go:embed templates/password_reset_html.tmpl
	verifyUserHTMLTMPL string
	//go:embed templates/password_reset_text.tmpl
	verifyUserTextTMPL string
	ContentMap         = mapContent()
)

type Content struct {
	html *htmlTmpl.Template
	text *textTmpl.Template
}

func mapContent() Contents {
	res := map[ContentType]*Content{
		ResetPassword: MustParseContent(ResetPassword, passwordResetHTMLTMPL, passwordResetTextTMPL),
		VerifyUser:    MustParseContent(ResetPassword, verifyUserHTMLTMPL, verifyUserTextTMPL),
	}
	return res
}

func parseContent(ct ContentType, html string, text string) (*Content, error) {
	t1, err := htmlTmpl.New(string(ct)).Parse(html)
	if err != nil {
		return nil, err
	}
	t2, err := textTmpl.New(string(ct)).Parse(text)
	if err != nil {
		return nil, err
	}
	return &Content{
		html: t1,
		text: t2,
	}, nil
}
func MustParseContent(ct ContentType, html string, text string) *Content {
	c, err := parseContent(ct, html, text)
	if err != nil {
		panic(err)
	}
	return c
}

func (cl Contents) GetContentByType(t ContentType) *Content {
	for k, v := range cl {
		if k == t {
			return v
		}
	}
	return nil
}

func (c *Content) Html() *htmlTmpl.Template {
	return c.html
}

func (c *Content) Text() *textTmpl.Template {
	return c.text
}
