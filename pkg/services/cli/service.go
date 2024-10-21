package cli

import "github.com/jacobbrewer1/league-manager/pkg/repositories/cli"

type Service interface {
	// Run is the entry point for the service
	Run() error
}

type service struct {
	// r is the repository
	r cli.Repository
}

func NewService(r cli.Repository) Service {
	return &service{
		r: r,
	}
}
