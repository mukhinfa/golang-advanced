package email

import (
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/mukhinfa/golang-advanced/3-validation-api/configs"
)

type Email struct{}

func (e *Email) Send(config configs.Config, to, subject, body string) error {
	email := email.NewEmail()
	email.To = []string{to}
	email.Subject = subject
	email.Text = []byte(body)
	if err := email.Send(config.Address, smtp.PlainAuth("", config.Email, config.Password, "smtp.gmail.com")); err != nil {
		return err
	}
	return nil
}

func New() *Email {
	return &Email{}
}
