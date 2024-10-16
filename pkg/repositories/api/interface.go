package api

import (
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	"github.com/Jacobbrewer1/pagefilter"
)

type Repository interface {
	// GetPlayers gets a list of players.
	GetPlayers(details pagefilter.PaginatorDetails, filters *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error)
}
