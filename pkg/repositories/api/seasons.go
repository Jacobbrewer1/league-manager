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
	// ErrSeasonNotFound is returned when a season is not found
	ErrSeasonNotFound = errors.New("season not found")

	// ErrDuplicateSeason is returned when a season already exists
	ErrDuplicateSeason = errors.New("season already exists")
)

func (r *repository) GetSeasons(details *pagefilter.PaginatorDetails, filters *GetSeasonsFilters) (*pagefilter.PaginatedResponse[models.Season], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_seasons"))
	defer t.ObserveDuration()

	mf := r.getSeasonsFilters(filters)

	pg := pagefilter.NewPaginator(r.db, "season", "id", mf)

	if err := pg.SetDetails(details, "id", "name"); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrSeasonNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	pvt, err := pg.Pivot()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrSeasonNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	items := make([]*models.Season, 0)
	if err := pg.Retrieve(pvt, &items); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrSeasonNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	var total int64 = 0
	if err := pg.Counts(&total); err != nil {
		return nil, fmt.Errorf("set paginator details: %w", err)
	}

	return &pagefilter.PaginatedResponse[models.Season]{
		Items: items,
		Total: total,
	}, nil
}

func (r *repository) getSeasonsFilters(got *GetSeasonsFilters) *pagefilter.MultiFilter {
	mf := pagefilter.NewMultiFilter()
	if got == nil {
		return mf
	}

	if got.Name != "" {
		mf.Add(filters.NewNameLike(got.Name))
	}

	return mf
}

func (r *repository) CreateSeason(season *models.Season) error {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("create_season"))
	defer t.ObserveDuration()

	seasonByName, err := models.SeasonByName(r.db, season.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("get season by name: %w", err)
	} else if seasonByName != nil || err == nil {
		return ErrDuplicateSeason
	}

	if err := season.Insert(r.db); err != nil {
		return fmt.Errorf("insert season: %w", err)
	}

	return nil
}

func (r *repository) GetSeason(id int64) (*models.Season, error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_season"))
	defer t.ObserveDuration()

	season, err := models.SeasonById(r.db, int(id))
	if err != nil {
		return nil, fmt.Errorf("get season by ID: %w", err)
	}

	return season, nil
}

func (r *repository) UpdateSeason(id int64, season *models.Season) error {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("update_season"))
	defer t.ObserveDuration()

	season.Id = int(id)

	if err := season.Update(r.db); err != nil {
		return fmt.Errorf("update season: %w", err)
	}

	return nil
}
