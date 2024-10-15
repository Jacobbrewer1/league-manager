package api

import (
	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
)

type service struct{}

// NewService creates a new service.
func NewService() api.ServerInterface {
	return &service{}
}
