package api

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Jacobbrewer1/league-manager/pkg/models"
	"github.com/Jacobbrewer1/league-manager/pkg/repositories/api/filters"
	"github.com/Jacobbrewer1/pagefilter"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ErrTeamNotFound is returned when a team is not found
	ErrTeamNotFound = errors.New("team not found")
)

func (r *repository) GetTeams(details *pagefilter.PaginatorDetails, filters *GetTeamsFilters) (*pagefilter.PaginatedResponse[models.Team], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_teams"))
	defer t.ObserveDuration()

	mf := r.getTeamsFilters(filters)
	pg := pagefilter.NewPaginator(r.db, "team", "id", mf)

	if err := pg.SetDetails(details, "id", "name"); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrTeamNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	pvt, err := pg.Pivot()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrTeamNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	items := make([]*models.Team, 0)
	err = pg.Retrieve(pvt, &items)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrTeamNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	var total int64 = 0
	err = pg.Counts(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	return &pagefilter.PaginatedResponse[models.Team]{
		Items: items,
		Total: total,
	}, nil
}

func (r *repository) getTeamsFilters(got *GetTeamsFilters) *pagefilter.MultiFilter {
	mf := pagefilter.NewMultiFilter()
	if got == nil {
		return mf
	}

	if got.Name != "" {
		mf.Add(filters.NewNameLike(got.Name))
	}

	return mf
}
