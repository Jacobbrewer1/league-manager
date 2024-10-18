// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	externalRef0 "github.com/Jacobbrewer1/pagefilter/common"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Game defines the model for game.
type Game struct {
	AwayTeam    *GamePartnership `json:"away_team,omitempty"`
	HomeTeam    *GamePartnership `json:"home_team,omitempty"`
	Id          *int64           `json:"id,omitempty"`
	MatchDate   *string          `json:"match_date,omitempty"`
	Season      *Season          `json:"season,omitempty"`
	WinningTeam *WinningTeam     `json:"winning_team,omitempty"`
}

// GamePartnership defines the model for game_partnership.
type GamePartnership struct {
	Partnership *Partnership `json:"partnership,omitempty"`
	Scores      *Scores      `json:"scores,omitempty"`
}

// GamesResponse defines the model for games_response.
type GamesResponse struct {
	Games []Game `json:"games"`
	Total int64  `json:"total"`
}

// Partnership defines the model for partnership.
type Partnership struct {
	PlayerA *Player `json:"player_a,omitempty"`
	PlayerB *Player `json:"player_b,omitempty"`
	Team    *Team   `json:"team,omitempty"`
}

// Player defines the model for player.
type Player struct {
	DateOfBirth *openapi_types.Date  `json:"date_of_birth,omitempty"`
	Email       *openapi_types.Email `json:"email,omitempty"`
	FirstName   *string              `json:"first_name,omitempty"`
	Id          *int64               `json:"id,omitempty"`
	LastName    *string              `json:"last_name,omitempty"`
}

// PlayersResponse defines the model for players_response.
type PlayersResponse struct {
	Players []Player `json:"players"`
	Total   int64    `json:"total"`
}

// Scores defines the model for scores.
type Scores struct {
	FirstSet  int64  `json:"first_set"`
	SecondSet int64  `json:"second_set"`
	ThirdSet  *int64 `json:"third_set,omitempty"`
}

// Season defines the model for season.
type Season struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// SeasonsResponse defines the model for seasons_response.
type SeasonsResponse struct {
	Seasons []Season `json:"seasons"`
	Total   int64    `json:"total"`
}

// Team defines the model for team.
type Team struct {
	ContactEmail *openapi_types.Email `json:"contact_email,omitempty"`
	ContactPhone *string              `json:"contact_phone,omitempty"`
	Id           *int64               `json:"id,omitempty"`
	Name         *string              `json:"name,omitempty"`
}

// TeamsResponse defines the model for teams_response.
type TeamsResponse struct {
	Teams []Team `json:"teams"`
	Total int64  `json:"total"`
}

// WinningTeam defines the model for winning_team.
type WinningTeam = string

// List of WinningTeam
const (
	WinningTeam_away WinningTeam = "away"
	WinningTeam_home WinningTeam = "home"
)

// QueryDate defines the model for query_date.
type QueryDate = openapi_types.Date

// QueryDateMax defines the model for query_date_max.
type QueryDateMax = openapi_types.Date

// QueryDateMin defines the model for query_date_min.
type QueryDateMin = openapi_types.Date

// QueryName defines the model for query_name.
type QueryName = string

// QuerySeason defines the model for query_season.
type QuerySeason = string

// QueryTeam defines the model for query_team.
type QueryTeam = string

// QueryYear defines the model for query_year.
type QueryYear = int64

// GetGamesParams defines parameters for GetGames.
type GetGamesParams struct {
	// Limit Report type
	Limit *externalRef0.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef0.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef0.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef0.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetGamesParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Date The date to filter by
	Date *QueryDate `form:"date,omitempty" json:"date,omitempty"`

	// DateMin The minimum date to filter by
	DateMin *QueryDateMin `form:"date_min,omitempty" json:"date_min,omitempty"`

	// DateMax The maximum date to filter by
	DateMax *QueryDateMax `form:"date_max,omitempty" json:"date_max,omitempty"`

	// Season The season to filter by
	Season *QuerySeason `form:"season,omitempty" json:"season,omitempty"`

	// Team The team to filter by
	Team *QueryTeam `form:"team,omitempty" json:"team,omitempty"`
}

// GetGamesParamsSortDir defines parameters for GetGames.
type GetGamesParamsSortDir string

// GetPlayersParams defines parameters for GetPlayers.
type GetPlayersParams struct {
	// Limit Report type
	Limit *externalRef0.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef0.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef0.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef0.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetPlayersParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Name The name to filter by
	Name *QueryName `form:"name,omitempty" json:"name,omitempty"`

	// Year The year of the season
	Year *QueryYear `form:"year,omitempty" json:"year,omitempty"`
}

// GetPlayersParamsSortDir defines parameters for GetPlayers.
type GetPlayersParamsSortDir string

// GetSeasonsParams defines parameters for GetSeasons.
type GetSeasonsParams struct {
	// Limit Report type
	Limit *externalRef0.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef0.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef0.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef0.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetSeasonsParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Name The name to filter by
	Name *QueryName `form:"name,omitempty" json:"name,omitempty"`
}

// GetSeasonsParamsSortDir defines parameters for GetSeasons.
type GetSeasonsParamsSortDir string

// GetTeamsParams defines parameters for GetTeams.
type GetTeamsParams struct {
	// Limit Report type
	Limit *externalRef0.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef0.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef0.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef0.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetTeamsParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Name The name to filter by
	Name *QueryName `form:"name,omitempty" json:"name,omitempty"`
}

// GetTeamsParamsSortDir defines parameters for GetTeams.
type GetTeamsParamsSortDir string

// CreateGameJSONRequestBody defines body for CreateGame for application/json ContentType.
type CreateGameJSONRequestBody = Game

// Temporary inclusion of type alias for backwards compatibility
type CreateGameJSONBody = Game

// CreatePlayerJSONRequestBody defines body for CreatePlayer for application/json ContentType.
type CreatePlayerJSONRequestBody = Player

// Temporary inclusion of type alias for backwards compatibility
type CreatePlayerJSONBody = Player

// UpdatePlayerJSONRequestBody defines body for UpdatePlayer for application/json ContentType.
type UpdatePlayerJSONRequestBody = Player

// Temporary inclusion of type alias for backwards compatibility
type UpdatePlayerJSONBody = Player

// CreateSeasonJSONRequestBody defines body for CreateSeason for application/json ContentType.
type CreateSeasonJSONRequestBody = Season

// Temporary inclusion of type alias for backwards compatibility
type CreateSeasonJSONBody = Season

// UpdateSeasonJSONRequestBody defines body for UpdateSeason for application/json ContentType.
type UpdateSeasonJSONRequestBody = Season

// Temporary inclusion of type alias for backwards compatibility
type UpdateSeasonJSONBody = Season

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = Team

// Temporary inclusion of type alias for backwards compatibility
type CreateTeamJSONBody = Team

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.
type UpdateTeamJSONRequestBody = Team

// Temporary inclusion of type alias for backwards compatibility
type UpdateTeamJSONBody = Team
