package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	mrand "math/rand"
	"sync"
	"time"
)

const sessionTTL = 5 * time.Minute

type JWT interface {
	Create(phone string) (string, error)
}

type session struct {
	Phone     string
	Code      int
	ExpiresAt time.Time
}

type authService struct {
	JWT
	mu       sync.Mutex
	sessions map[string]session
}

func NewAuthService(jwt JWT) *authService {
	return &authService{
		JWT:      jwt,
		sessions: make(map[string]session),
	}
}

func (s *authService) SendCode(phone string) (string, error) {
	code := mrand.Intn(9000) + 1000

	// === MOCK SMS ===
	log.Printf("[MOCK SMS] phone=%s code=%d", phone, code)
	// ================

	sessionID, err := newSessionID()
	if err != nil {
		return "", errors.New(ErrSendSMS)
	}

	s.mu.Lock()
	s.sessions[sessionID] = session{
		Phone:     phone,
		Code:      code,
		ExpiresAt: time.Now().Add(sessionTTL),
	}
	s.mu.Unlock()

	return sessionID, nil
}

func (s *authService) VerifyCode(sessionID string, code int) (string, error) {
	s.mu.Lock()
	sess, ok := s.sessions[sessionID]
	if !ok {
		s.mu.Unlock()
		return "", errors.New(ErrSessionNotFound)
	}
	if time.Now().After(sess.ExpiresAt) {
		delete(s.sessions, sessionID)
		s.mu.Unlock()
		return "", errors.New(ErrSessionExpired)
	}
	if sess.Code != code {
		s.mu.Unlock()
		return "", errors.New(ErrWrongCode)
	}
	delete(s.sessions, sessionID)
	s.mu.Unlock()

	return s.generateToken(sess.Phone)
}

func (s *authService) generateToken(phone string) (string, error) {
	token, err := s.JWT.Create(phone)
	if err != nil {
		log.Printf("failed to create JWT: %v", err)
		return "", errors.New(ErrGenerateToken)
	}
	return token, nil
}

func newSessionID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
