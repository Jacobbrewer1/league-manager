package filters

import "github.com/Jacobbrewer1/pagefilter"

type matchTeamNameLike struct {
	teamName string
}

// NewMatchTeamNameLike creates a new match team name like.
func NewMatchTeamNameLike(teamName string) pagefilter.Wherer {
	return &matchTeamNameLike{
		teamName: teamName,
	}
}

// Where returns the where clause for the match team name like.
func (m *matchTeamNameLike) Where() (string, []interface{}) {
	return "tm.name LIKE ?", []interface{}{"%" + m.teamName + "%"}
}
