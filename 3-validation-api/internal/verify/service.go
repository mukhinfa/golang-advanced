package verify

import (
	"github.com/mukhinfa/golang-advanced/3-validation-api/configs"
	"github.com/mukhinfa/golang-advanced/3-validation-api/pkg/email"
	"github.com/mukhinfa/golang-advanced/3-validation-api/pkg/utils"
)

type Storage interface {
	Save(email, hash string) error
	Find(hash string) (string, error)
	Delete(hash string) error
}

type SendEmail interface {
	Send(c configs.Config, to, subject, body string) error
}

const domain = "http://localhost:8081"
const verifyPath = "/verify/"
const emailSubject = "Verify Email"

type VerifyService struct {
	storage Storage
	config  configs.Config
	email   SendEmail
}

func NewVerifyService(s Storage, c configs.Config) *VerifyService {
	return &VerifyService{
		storage: s,
		config:  c,
		email:   email.New(),
	}
}

func (v *VerifyService) SendEmail(emailAddr string) error {
	hash, err := utils.GenerateHash()
	if err != nil {
		return err
	}

	if err := v.email.Send(v.config,
		emailAddr,
		emailSubject,
		domain+verifyPath+hash); err != nil {
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
