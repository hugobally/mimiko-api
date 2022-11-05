package authentication

import (
	"context"
	jwtlib "github.com/golang-jwt/jwt/v4"
	"github.com/hugobally/mimiko_api/internal/shared"
	"net/http"
)

type Middleware struct {
	*shared.Services

	Next http.Handler
}

func NewMiddleware(n http.Handler, s *shared.Services) *Middleware {
	m := &Middleware{s, n}
	return m
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Config.Env == "DEV" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	tokenStr := c.Value
	user, err := m.JwtUtil.ValidateAccessToken(tokenStr)
	if err != nil {
		m.Logger.Println("invalid access token received")
		if err == jwtlib.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	ctxWithUser := context.WithValue(r.Context(), userCtxKey, user)
	rWithContext := r.WithContext(ctxWithUser)

	m.Next.ServeHTTP(w, rWithContext)
	return
}
