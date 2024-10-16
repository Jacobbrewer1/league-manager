package api

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
