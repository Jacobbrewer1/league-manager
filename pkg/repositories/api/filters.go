package api

import openapi_types "github.com/oapi-codegen/runtime/types"

type GetPlayersFilters struct {
	// Name is the name of the player.
	Name string `json:"name"`
}

type GetTeamsFilters struct {
	// Name is the name of the team.
	Name string `json:"name"`
}

type GetSeasonsFilters struct {
	// Name is the name of the season.
	Name string `json:"name"`
}

type GetMatchesFilters struct {
	// Season is the season of the match.
	Season string `json:"season"`

	// Team is the team of the match.
	Team string `json:"team"`

	// Date is the date of the match.
	Date *openapi_types.Date `json:"date"`

	// DateFrom is the date from of the match.
	DateFrom *openapi_types.Date `json:"date_from"`

	// DateTo is the date to of the match.
	DateTo *openapi_types.Date `json:"date_to"`
}
