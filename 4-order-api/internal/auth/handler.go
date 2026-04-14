package auth

import (
	"errors"
	"net/http"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/req"
	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/res"
)

type AuthHandlerDeps struct {
	Service *authService
}

func NewHandlerDeps(service *authService) AuthHandlerDeps {
	return AuthHandlerDeps{
		Service: service,
	}
}

type authHandler struct {
	service *authService
}

func NewHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		service: deps.Service,
	}
	r.HandleFunc("POST /auth/phone", handler.sendCode())
	r.HandleFunc("POST /auth/verify", handler.verify())
}

func (h *authHandler) sendCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionID, err := h.service.SendCode(body.Phone)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		res.JSON(w, http.StatusOK, PhoneResponse{SessionID: sessionID})
	}
}

func (h *authHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := h.service.VerifyCode(body.SessionID, body.Code)
		if err != nil {
			switch {
			case errors.Is(err, errors.New(ErrGenerateToken)):
				http.Error(w, "internal server error", http.StatusInternalServerError)
			default:
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			return
		}
		res.JSON(w, http.StatusAccepted, VerifyResponse{Token: token})
	}
}
