package middleware

import (
	"context"
	"net/http"
	"strings"
)

type TokenParser interface {
	Parse(token string) (phone string, err error)
}

type ctxKey string

const PhoneCtxKey ctxKey = "phone"

func IsAuthed(parser TokenParser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")
			phone, err := parser.Parse(token)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), PhoneCtxKey, phone)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetPhone(r *http.Request) *string {
	phone, ok := r.Context().Value(PhoneCtxKey).(string)
	if !ok {
		return nil
	}
	return &phone
}
