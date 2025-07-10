package verify

import (
	"net/http"

	"github.com/mukhinfa/golang-advanced/3-validation-api/internal/configs"
)

type VerifyHandler struct {
	config configs.Config
}

func NewVerifyHandler(router *http.ServeMux) {
	handler := &VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (h VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (h VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
