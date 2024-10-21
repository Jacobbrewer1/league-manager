package cli

import (
	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
)

type Repository interface {
	// GetPlayers gets a list of players.
	GetPlayers(params *api.GetPlayersParams) (*api.PlayersResponse, error)
}
