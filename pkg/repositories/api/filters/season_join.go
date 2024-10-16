package filters

import "github.com/Jacobbrewer1/pagefilter"

type seasonJoin struct{}

// NewSeasonJoin creates a new season join.
func NewSeasonJoin() pagefilter.Joiner {
	return &seasonJoin{}
}

// Join joins the season join.
func (s *seasonJoin) Join() (string, []interface{}) {
	return "JOIN season s ON s.id = t.season_id", nil
}
