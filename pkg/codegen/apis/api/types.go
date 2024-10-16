// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	externalRef1 "github.com/Jacobbrewer1/pagefilter/common"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

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

// QueryName defines the model for query_name.
type QueryName = string

// QueryYear defines the model for query_year.
type QueryYear = int64

// GetPlayersParams defines parameters for GetPlayers.
type GetPlayersParams struct {
	// Limit Report type
	Limit *externalRef1.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef1.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef1.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef1.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetPlayersParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Name The name to filter by
	Name *QueryName `form:"name,omitempty" json:"name,omitempty"`

	// Year The year of the season
	Year *QueryYear `form:"year,omitempty" json:"year,omitempty"`
}

// GetPlayersParamsSortDir defines parameters for GetPlayers.
type GetPlayersParamsSortDir string

// GetTeamsParams defines parameters for GetTeams.
type GetTeamsParams struct {
	// Limit Report type
	Limit *externalRef1.LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// LastVal Pagination details, last value of the sort column on the previous page.
	LastVal *externalRef1.LastValue `form:"last_val,omitempty" json:"last_val,omitempty"`

	// LastId Pagination details, last value of the id column on the previous page.
	LastId *externalRef1.LastId `form:"last_id,omitempty" json:"last_id,omitempty"`

	// SortBy Pagination details, sort column, if empty uses the id column.
	SortBy *externalRef1.SortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`

	// SortDir Pagination details, sorting order.
	SortDir *GetTeamsParamsSortDir `form:"sort_dir,omitempty" json:"sort_dir,omitempty"`

	// Name The name to filter by
	Name *QueryName `form:"name,omitempty" json:"name,omitempty"`
}

// GetTeamsParamsSortDir defines parameters for GetTeams.
type GetTeamsParamsSortDir string

// CreatePlayerJSONRequestBody defines body for CreatePlayer for application/json ContentType.
type CreatePlayerJSONRequestBody = Player

// Temporary inclusion of type alias for backwards compatibility
type CreatePlayerJSONBody = Player

// UpdatePlayerJSONRequestBody defines body for UpdatePlayer for application/json ContentType.
type UpdatePlayerJSONRequestBody = Player

// Temporary inclusion of type alias for backwards compatibility
type UpdatePlayerJSONBody = Player

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = Team

// Temporary inclusion of type alias for backwards compatibility
type CreateTeamJSONBody = Team
