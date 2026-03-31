package verify

import (
	"crypto/rand"
	"encoding/hex"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/mukhinfa/golang-advanced/3-validation-api/configs"
)

type Storage interface {
	Save(email, hash string) error
	Find(hash string) (string, error)
	Delete(hash string) error
}

const domain = "http://localhost:8081"

type VerifyService struct {
	storage Storage
	config  configs.Config
}

func NewVerifyService(s Storage, c configs.Config) *VerifyService {
	return &VerifyService{
		storage: s,
		config:  c,
	}
}

func (v *VerifyService) SendEmail(emailAddr string) error {
	e := email.NewEmail()
	e.To = []string{emailAddr}

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	hash := hex.EncodeToString(b)
	text := domain + "/verify/" + hash

	e.Text = []byte(text)
	e.From = v.config.Email
	e.Subject = "Email Verification"
	if err := e.Send(v.config.Address, smtp.PlainAuth("", v.config.Email, v.config.Password, "smtp.gmail.com")); err != nil {
		return err
	}

	if err := v.storage.Save(emailAddr, hash); err != nil {
		return err
	}
	return nil
}

func (v *VerifyService) VerifyEmail(hash string) (bool, error) {
	_, err := v.storage.Find(hash)
	if err != nil {
		return false, err
	}
	if err := v.storage.Delete(hash); err != nil {
		return false, err
	}
	return true, nil
}
