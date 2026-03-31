package verify

import (
	"log"
	"net/http"

	"github.com/mukhinfa/golang-advanced/3-validation-api/configs"
	"github.com/mukhinfa/golang-advanced/3-validation-api/pkg/req"
	"github.com/mukhinfa/golang-advanced/3-validation-api/pkg/res"
)

type VerifyHandler struct {
	config  configs.Config
	service *VerifyService
}

func NewVerifyHandler(r *http.ServeMux, c configs.Config, v *VerifyService) {
	handler := &VerifyHandler{
		config:  c,
		service: v,
	}
	r.HandleFunc("POST /send", handler.Send())
	r.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (h VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			log.Println("Error handling request body:", err)
			return
		}
		if err := h.service.SendEmail(body.Email); err != nil {
			res.JSONResponse(w, http.StatusInternalServerError, err)
			return
		}
		res.JSONResponse(w, http.StatusOK, "")
	}
}
func (h VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		url, err := h.service.VerifyEmail(hash)
		if err != nil {
			res.JSONResponse(w, http.StatusBadRequest, err)
			return
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
