package filters

import "github.com/Jacobbrewer1/pagefilter"

type firstnameLike struct {
	fnl string
}

// NewFirstnameLike creates a new instance of a firstname like filter.
func NewFirstnameLike(fn string) pagefilter.Wherer {
	return &firstnameLike{
		fnl: fn,
	}
}

// Where returns the where clause for the filter.
func (f *firstnameLike) Where() (string, []interface{}) {
	return "first_name LIKE ?", []interface{}{"%" + f.fnl + "%"}
}

func (f *firstnameLike) WhereType() pagefilter.WhereType {
	return pagefilter.WhereTypeOr
}
