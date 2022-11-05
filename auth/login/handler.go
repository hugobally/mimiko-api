package login

import (
	"github.com/hugobally/mimiko/backend/shared"
	"net/http"
)

type Handler struct{ *shared.Services }

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	//mux.HandleFunc("/login_spotify", h.LoginAuthCode)
	mux.HandleFunc("/login_sample_session", h.LoginSampleSession)
	mux.HandleFunc("/logout", h.Logout)
}

func NewHandler(s *shared.Services) *Handler {
	return &Handler{s}
}
