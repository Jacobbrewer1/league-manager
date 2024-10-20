package filters

import (
	"strings"

	"github.com/jacobbrewer1/pagefilter"
)

type partnershipJoin struct{}

// NewPartnershipJoin creates a new partnership join.
func NewPartnershipJoin() pagefilter.Joiner {
	return &partnershipJoin{}
}

// Join joins the partnership join.
func (p *partnershipJoin) Join() (string, []interface{}) {
	str := new(strings.Builder)
	str.WriteString("JOIN partnership hp ON hp.id = t.home_partners_id")
	str.WriteString("\n")
	str.WriteString("JOIN partnership ap ON ap.id = t.away_partners_id")
	return str.String(), nil
}
