package static

import (
	"fmt"
	"github.com/hugobally/mimiko/backend/shared"
	"net/http"
)

type Handler struct{ *shared.Services }

func NewHandler(s *shared.Services) *Handler {
	return &Handler{s}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	staticPath := h.Config.Server.StaticPath
	indexPath := fmt.Sprintf("%v/%v", staticPath, "index.html")
	mux.Handle("/", WrapNotFound(http.FileServer(http.Dir(staticPath)), indexPath))
}

// Vue Frontend requires 404's on static assets to return index.html instead
func WrapNotFound(next http.Handler, indexPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cw := &CustomResponseWriter{ResponseWriter: w}
		next.ServeHTTP(cw, r)
		if cw.status == http.StatusNotFound {
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			http.ServeFile(w, r, indexPath)
		}
	}
}

type CustomResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *CustomResponseWriter) WriteHeader(status int) {
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	} else {
		w.status = status
	}
}

func (w *CustomResponseWriter) Write(p []byte) (int, error) {
	if w.status != http.StatusNotFound {
		return w.ResponseWriter.Write(p)
	} else {
		return len(p), nil
	}
}
