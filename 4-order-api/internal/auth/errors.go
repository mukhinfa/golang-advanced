package auth

const (
	ErrSessionNotFound = "session not found"
	ErrSessionExpired  = "session expired"
	ErrWrongCode       = "wrong code"
	ErrSendSMS         = "can't send sms"
	ErrGenerateToken   = "can't generate token"
)
