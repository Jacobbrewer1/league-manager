package filters

import "github.com/Jacobbrewer1/pagefilter"

type lastnameLike struct {
	lnl string
}

// NewLastnameLike creates a new instance of a lastname like filter.
func NewLastnameLike(ln string) pagefilter.Wherer {
	return &lastnameLike{
		lnl: ln,
	}
}

// Where returns the where clause for the filter.
func (f *lastnameLike) Where() (string, []interface{}) {
	return "last_name LIKE ?", []interface{}{f.lnl}
}

func (f *lastnameLike) WhereType() pagefilter.WhereType {
	return pagefilter.WhereTypeOr
}
