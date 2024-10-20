package filters

import (
	"github.com/jacobbrewer1/pagefilter"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type dateFrom struct {
	date *openapi_types.Date
}

// NewDateFrom creates a new date from.
func NewDateFrom(date *openapi_types.Date) pagefilter.Wherer {
	return &dateFrom{
		date: date,
	}
}

// Where returns the where clause for the date from.
func (d *dateFrom) Where() (string, []interface{}) {
	return "t.match_date >= ?", []interface{}{d.date.Time}
}
