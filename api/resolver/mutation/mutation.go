package mutation

import (
	"github.com/hugobally/mimiko/backend/api/permission"
	"github.com/hugobally/mimiko/backend/shared"
)

type Resolver struct {
	*shared.Services

	Permission *permission.Client
}

func New(s *shared.Services, p *permission.Client) *Resolver {
	return &Resolver{s, p}
}
