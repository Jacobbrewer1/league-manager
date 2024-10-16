package filters

import (
	"strings"

	"github.com/Jacobbrewer1/pagefilter"
)

type teamJoin struct{}

// NewTeamJoin creates a new team join.
func NewTeamJoin() pagefilter.Joiner {
	return &teamJoin{}
}

// Join joins the team join.
func (t *teamJoin) Join() (string, []interface{}) {
	str := new(strings.Builder)
	str.WriteString("JOIN team tm ON tm.id = hp.team_id")
	str.WriteString("\n")
	str.WriteString("JOIN team ta ON ta.id = ap.team_id")
	return str.String(), nil
}
