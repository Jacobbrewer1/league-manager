package filters

import "github.com/Jacobbrewer1/pagefilter"

type nameLike struct {
	nl string
}

// NewNameLike creates a new name like filter.
func NewNameLike(nl string) pagefilter.Wherer {
	return &nameLike{
		nl: nl,
	}
}

// Where returns the where clause for the filter.
func (n *nameLike) Where() (string, []interface{}) {
	return "name like ?", []interface{}{n.nl}
}