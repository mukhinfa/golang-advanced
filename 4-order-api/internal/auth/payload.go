package auth

type PhoneRequest struct {
	Phone string `json:"phone" validate:"required,len=11,numeric,startswith=8"`
}

type PhoneResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyRequest struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
