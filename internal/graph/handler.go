package graph

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hugobally/mimiko_api/internal/authentication"
	"github.com/hugobally/mimiko_api/internal/graph/generated"
	"github.com/hugobally/mimiko_api/internal/graph/resolver"
	"github.com/hugobally/mimiko_api/internal/shared"
	"net/http"
)

type Handler struct{ *shared.Services }

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	r := resolver.NewRoot(h.Services)
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	//Log GraphQL query content
	if h.Config.Env == "DEV" {
		gqlHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			oc := graphql.GetOperationContext(ctx)
			fmt.Printf("%s\n%s\n%s\n", oc.OperationName, oc.RawQuery, oc.Variables)
			return next(ctx)
		})
	}

	mux.Handle("/graphql", authentication.NewMiddleware(gqlHandler, h.Services))
}

func NewHandler(s *shared.Services) *Handler {
	return &Handler{s}
}
