package filters

import (
	"github.com/jacobbrewer1/pagefilter"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type dateExact struct {
	date *openapi_types.Date
}

// NewDateExact creates a new date exact.
func NewDateExact(date *openapi_types.Date) pagefilter.Wherer {
	return &dateExact{
		date: date,
	}
}

// Where returns the where clause for the date exact.
func (d *dateExact) Where() (string, []interface{}) {
	// Set t1 to be that start of the day and t2 to be the end of the day.
	t1 := d.date.Time
	t2 := t1.AddDate(0, 0, 1).Add(-1)
	return "t.match_date between ? and ?", []interface{}{t1, t2}
}
