package api

import (
	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
	repo "github.com/jacobbrewer1/league-manager/pkg/repositories/api"
)

type service struct {
	// r is the repository.
	r repo.Repository
}

// NewService creates a new service.
func NewService(r repo.Repository) api.ServerInterface {
	return &service{
		r: r,
	}
}
