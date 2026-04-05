package middleware

import (
	"fmt"
	"net/http"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Unauthorized")
			next.ServeHTTP(w, r)
		}
		token := authHeader[len("Bearer "):]
		fmt.Printf("token: %s\n", token)
		next.ServeHTTP(w, r)
	})
}
