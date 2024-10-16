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
	// ErrPlayerNotFound is returned when a player is not found
	ErrPlayerNotFound = errors.New("player not found")
)

func (r *repository) GetPlayers(details *pagefilter.PaginatorDetails, filters *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_players"))
	defer t.ObserveDuration()

	mf := r.getPlayersFilters(filters)
	pg := pagefilter.NewPaginator(r.db, "player", "id", mf)

	if err := pg.SetDetails(details, "id", "name"); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrPlayerNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	pvt, err := pg.Pivot()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrPlayerNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	items := make([]*models.Player, 0)
	err = pg.Retrieve(pvt, &items)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrPlayerNotFound
		default:
			return nil, fmt.Errorf("set paginator details: %w", err)
		}
	}

	var total int64 = 0
	err = pg.Counts(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	return &pagefilter.PaginatedResponse[models.Player]{
		Items: items,
		Total: total,
	}, nil
}

func (r *repository) getPlayersFilters(got *GetPlayersFilters) pagefilter.Filter {
	mf := pagefilter.NewMultiFilter()
	if got == nil {
		return mf
	}

	if got.Name != "" {
		mf.Add(filters.NewNameLike(got.Name))
	}

	return mf
}
