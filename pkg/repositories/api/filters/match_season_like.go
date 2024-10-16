package filters

import "github.com/Jacobbrewer1/pagefilter"

type matchSeasonLike struct {
	season string
}

// NewMatchSeasonLike creates a new match season like.
func NewMatchSeasonLike(season string) pagefilter.Wherer {
	return &matchSeasonLike{
		season: season,
	}
}

// Where returns the where clause for the match season like.
func (m *matchSeasonLike) Where() (string, []interface{}) {
	return "s.name LIKE ?", []interface{}{"%" + m.season + "%"}
}
