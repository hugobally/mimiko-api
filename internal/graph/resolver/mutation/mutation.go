package mutation

import (
	"github.com/hugobally/mimiko_api/internal/shared"
)

type MutationResolver struct {
	*shared.Services
}

func NewMutationResolver(s *shared.Services) *MutationResolver {
	return &MutationResolver{s}
}
