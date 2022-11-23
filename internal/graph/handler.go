package graph

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hugobally/mimiko_api/internal/authentication"
	"github.com/hugobally/mimiko_api/internal/graph/generated"
	"github.com/hugobally/mimiko_api/internal/graph/resolver"
	"github.com/hugobally/mimiko_api/internal/shared"
	"net/http"
)

type Handler struct{ *shared.Services }

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	r := resolver.NewRoot(h.Services)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	if h.Config.Env == "DEV" {
		srv.Use(extension.Introspection{})
		mux.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	}

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		fmt.Printf("%s\n%s\n%s\n", oc.OperationName, oc.RawQuery, oc.Variables)
		return next(ctx)
	})

	mux.Handle("/graphql", authentication.NewMiddleware(srv, h.Services))
}

func NewHandler(s *shared.Services) *Handler {
	return &Handler{s}
}
