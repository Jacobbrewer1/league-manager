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

	// CreateTeam creates a team.
	CreateTeam(team *models.Team) error

	// GetTeam gets a team by ID.
	GetTeam(id int64) (*models.Team, error)

	// UpdateTeam updates a team.
	UpdateTeam(id int64, team *models.Team) error

	// GetSeasons gets a list of seasons.
	GetSeasons(details *pagefilter.PaginatorDetails, filters *GetSeasonsFilters) (*pagefilter.PaginatedResponse[models.Season], error)

	// CreateSeason creates a season.
	CreateSeason(season *models.Season) error

	// GetSeason gets a season by ID.
	GetSeason(id int64) (*models.Season, error)

	// UpdateSeason updates a season.
	UpdateSeason(id int64, season *models.Season) error
}
