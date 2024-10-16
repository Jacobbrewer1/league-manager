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
	// ErrMatchNotFound is returned when a match is not found
	ErrMatchNotFound = errors.New("match not found")

	// ErrDuplicateMatch is returned when a match already exists
	ErrDuplicateMatch = errors.New("match already exists")
)

func (r *repository) GetMatches(details *pagefilter.PaginatorDetails, filters *GetMatchesFilters) (*pagefilter.PaginatedResponse[models.Match], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_matches"))
	defer t.ObserveDuration()

	mf := r.getMatchesFilters(filters)

	pg := pagefilter.NewPaginator(r.db, "`match`", "id", mf)

	if err := pg.SetDetails(details, "id", "name"); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrMatchNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	pvt, err := pg.Pivot()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrMatchNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	items := make([]*models.Match, 0)
	if err := pg.Retrieve(pvt, &items); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrMatchNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	var total int64 = 0
	if err := pg.Counts(&total); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrMatchNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	return &pagefilter.PaginatedResponse[models.Match]{
		Items: items,
		Total: total,
	}, nil
}

func (r *repository) getMatchesFilters(got *GetMatchesFilters) *pagefilter.MultiFilter {
	mf := pagefilter.NewMultiFilter()
	if got == nil {
		return mf
	}

	if got.Season != "" {
		mf.Add(filters.NewSeasonJoin())
		mf.Add(filters.NewMatchSeasonLike(got.Season))
	}

	if got.Team != "" {
		mf.Add(filters.NewPartnershipJoin())
		mf.Add(filters.NewTeamJoin())
		mf.Add(filters.NewMatchTeamNameLike(got.Team))
	}

	if got.Date != nil {
		mf.Add(filters.NewDateExact(got.Date))
	}

	if got.DateFrom != nil {
		mf.Add(filters.NewDateFrom(got.DateFrom))
	}

	if got.DateTo != nil {
		mf.Add(filters.NewDateTo(got.DateTo))
	}

	return mf
}
