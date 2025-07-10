package verify

import (
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/mukhinfa/golang-advanced/3-validation-api/internal/configs"
)

type VerifyHandler struct {
	config configs.Config
}

func NewVerifyHandler(r *http.ServeMux, c configs.Config) {
	handler := &VerifyHandler{
		config: c,
	}
	r.HandleFunc("POST /send", handler.Send())
	r.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (h VerifyHandler) Send() http.HandlerFunc {
	e := email.NewEmail()
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (h VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
