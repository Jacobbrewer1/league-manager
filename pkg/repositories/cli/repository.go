package cli

import "github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"

type repository struct {
	client api.ClientWithResponsesInterface
}

func NewRepository(client api.ClientWithResponsesInterface) Repository {
	return &repository{
		client: client,
	}
}
