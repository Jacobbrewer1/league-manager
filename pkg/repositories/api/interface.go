package api

import (
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	"github.com/Jacobbrewer1/pagefilter"
)

type Repository interface {
	// GetPlayers gets a list of players.
	GetPlayers(details *pagefilter.PaginatorDetails, filters *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error)

	// CreatePlayer creates a player.
	CreatePlayer(player *models.Player) error

	// GetPlayer gets a player by ID.
	GetPlayer(id int64) (*models.Player, error)

	// UpdatePlayer updates a player.
	UpdatePlayer(id int64, player *models.Player) error

	// GetTeams gets a list of teams.
	GetTeams(details *pagefilter.PaginatorDetails, filters *GetTeamsFilters) (*pagefilter.PaginatedResponse[models.Team], error)
}
