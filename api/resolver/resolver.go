package resolver

import (
	"github.com/hugobally/mimiko/backend/api/gqlgen"
	"github.com/hugobally/mimiko/backend/api/permission"
	"github.com/hugobally/mimiko/backend/api/resolver/mutation"
	"github.com/hugobally/mimiko/backend/api/resolver/query"
	"github.com/hugobally/mimiko/backend/shared"
)

type Resolver struct {
	*shared.Services

	Permission *permission.Client
}

func New(s *shared.Services, p *permission.Client) *Resolver {
	return &Resolver{s, p}
}

//

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return mutation.New(r.Services, r.Permission)
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return query.NewResolver(r.Services, r.Permission)
}

//

func (r *Resolver) Map() gqlgen.MapResolver {
	return &mapResolver{r}
}

func (r *Resolver) User() gqlgen.UserResolver {
	return &userResolver{r}
}
