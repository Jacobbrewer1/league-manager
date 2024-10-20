package filters

import (
	"github.com/jacobbrewer1/pagefilter"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type dateTo struct {
	date *openapi_types.Date
}

// NewDateTo creates a new date to.
func NewDateTo(date *openapi_types.Date) pagefilter.Wherer {
	return &dateTo{
		date: date,
	}
}

// Where returns the where clause for the date to.
func (d *dateTo) Where() (string, []interface{}) {
	return "t.match_date <= ?", []interface{}{d.date.Time}
}
