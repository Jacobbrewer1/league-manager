package api

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Jacobbrewer1/league-manager/pkg/models"
	"github.com/Jacobbrewer1/league-manager/pkg/repositories/api/filters"
	"github.com/Jacobbrewer1/pagefilter"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ErrPlayerNotFound is returned when a player is not found
	ErrPlayerNotFound = errors.New("player not found")

	// ErrDuplicatePlayer is returned when a player already exists
	ErrDuplicatePlayer = errors.New("player already exists")
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
		mf.Add(filters.NewFirstnameLike(got.Name))
		mf.Add(filters.NewLastnameLike(got.Name))
	}

	return mf
}

func (r *repository) CreatePlayer(player *models.Player) error {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("create_player"))
	defer t.ObserveDuration()

	playerByEmail, err := models.PlayerByEmail(r.db, player.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("get player by email: %w", err)
	} else if playerByEmail != nil || err == nil {
		return ErrDuplicatePlayer
	}

	player.UpdatedAt = time.Now().UTC()
	player.Email = strings.ToLower(player.Email)

	if err := player.InsertWithUpdate(r.db); err != nil {
		return fmt.Errorf("insert player: %w", err)
	}

	return nil
}

func (r *repository) GetPlayer(id int64) (*models.Player, error) {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("get_player"))
	defer t.ObserveDuration()

	player, err := models.PlayerById(r.db, int(id))
	if err != nil {
		return nil, fmt.Errorf("get player by ID: %w", err)
	}

	return player, nil
}

func (r *repository) UpdatePlayer(id int64, player *models.Player) error {
	t := prometheus.NewTimer(models.DatabaseLatency.WithLabelValues("update_player"))
	defer t.ObserveDuration()

	player.Id = int(id)
	player.UpdatedAt = time.Now().UTC()

	if err := player.Update(r.db); err != nil {
		return fmt.Errorf("update player: %w", err)
	}

	return nil
}
