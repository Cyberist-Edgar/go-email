package email

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type Email struct {
	emailHeader
	host string
	body string
	auth smtp.Auth
}

type emailHeader struct {
	from        string
	to          []string
	subject     string
	contentType string
}

func (e *Email) SetTo(to ...string) {
	e.to = to
}

func (e *Email) SetSubject(subject string) {
	e.subject = subject
}

func (e *Email) SetAuth(username, password string) {
	// 发送者应该和username一致
	e.from = username
	e.auth = smtp.PlainAuth("", username, password, strings.Split(e.host, ":")[0])
}

func (e *Email) SetContentType(typ string) error {
	if typ != "text/plain" && typ != "text/html" {
		return errors.New("invalid content type, must be text/plain or text/html")
	}
	e.contentType = typ
	return nil
}

func (e *Email) SetHost(host string) error {
	h := strings.Split(host, ":")
	if len(h) == 1 {
		return errors.New("missing port")
	}
	e.host = host
	return nil
}

func (e *Email) SetBody(body string) {
	e.body = body
}

func (e *Email) SendMail() error {
	sendTo := strings.Join(e.to, ",")
	msg := fmt.Sprintf("From: %s> \r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n %s",
		e.from,
		sendTo,
		e.subject,
		e.contentType,
		e.body,
	)
	fmt.Println(msg)
	return smtp.SendMail(e.host, e.auth, e.from, e.to, []byte(msg))
}

func New(subject, contentType string) *Email {
	return &Email{
		emailHeader: emailHeader{
			subject:     subject,
			contentType: contentType,
		}}
}
