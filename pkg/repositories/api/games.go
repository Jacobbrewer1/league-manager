package api

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jacobbrewer1/league-manager/pkg/models"
	"github.com/jacobbrewer1/league-manager/pkg/repositories/api/filters"
	"github.com/jacobbrewer1/pagefilter"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ErrMatchNotFound is returned when a match is not found
	ErrMatchNotFound = errors.New("match not found")

	// ErrDuplicateMatch is returned when a match already exists
	ErrDuplicateMatch = errors.New("match already exists")
)

func (r *repository) GetGames(details *pagefilter.PaginatorDetails, filters *GetMatchesFilters) (*pagefilter.PaginatedResponse[models.Game], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_matches"))
	defer t.ObserveDuration()

	mf := r.getMatchesFilters(filters)

	pg := pagefilter.NewPaginator(r.db, "game", "id", mf)

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

	items := make([]*models.Game, 0)
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

	return &pagefilter.PaginatedResponse[models.Game]{
		Items: items,
		Total: total,
	}, nil
}

func (r *repository) GetGameDetails(details *pagefilter.PaginatorDetails, gameID int64) (*pagefilter.PaginatedResponse[models.Game], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_matches"))
	defer t.ObserveDuration()

	return nil, nil
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

func (r *repository) CreateMatch(match *models.Game) error {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("create_match"))
	defer t.ObserveDuration()

	if err := match.Insert(r.db); err != nil {
		return fmt.Errorf("insert match: %w", err)
	}

	return nil
}

func (r *repository) GetGame(id int64) (*models.Game, error) {
	g, err := models.GameById(r.db, int(id))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrMatchNotFound
		default:
			return nil, fmt.Errorf("get match: %w", err)
		}
	}

	return g, nil
}
