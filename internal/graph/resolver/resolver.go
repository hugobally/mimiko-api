package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/hugobally/mimiko_api/internal/graph/generated"
	"github.com/hugobally/mimiko_api/internal/graph/resolver/mutation"
	"github.com/hugobally/mimiko_api/internal/graph/resolver/query"
	"github.com/hugobally/mimiko_api/internal/shared"
)

type Root struct {
	*shared.Services
}

func NewRoot(s *shared.Services) *Root {
	return &Root{s}
}

func (r *Root) Mutation() generated.MutationResolver {
	return mutation.NewMutationResolver(r.Services)
}

func (r *Root) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.Services)
}

func (r *Root) Map() generated.MapResolver {
	return &MapResolver{r}
}

func (r *Root) User() generated.UserResolver {
	return &UserResolver{r}
}
