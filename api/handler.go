package api

import (
	"github.com/hugobally/mimiko/backend/api/permission"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/shared"
	"net/http"

	"github.com/hugobally/mimiko/backend/api/gqlgen"
	"github.com/hugobally/mimiko/backend/api/resolver"

	"github.com/99designs/gqlgen/handler"
)

type Handler struct{ *shared.Services }

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	r := resolver.New(h.Services, permission.NewClient(h.Services.Prisma))

	gqlHandler := handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: r}))

	mux.Handle("/graphql", auth.NewMiddleware(gqlHandler, h.Services))
}

func NewHandler(s *shared.Services) *Handler {
	return &Handler{s}
}
