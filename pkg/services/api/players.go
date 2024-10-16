package api

import (
	"errors"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/Jacobbrewer1/league-manager/pkg/logging"
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	repo "github.com/Jacobbrewer1/league-manager/pkg/repositories/api"
	"github.com/Jacobbrewer1/league-manager/pkg/utils"
	"github.com/Jacobbrewer1/pagefilter"
	"github.com/Jacobbrewer1/pagefilter/common"
	"github.com/Jacobbrewer1/patcher"
	"github.com/Jacobbrewer1/uhttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (s *service) GetPlayers(w http.ResponseWriter, r *http.Request, params api.GetPlayersParams) {
	l := logging.LoggerFromRequest(r)

	sortDir := new(common.SortDirection)
	if params.SortDir != nil {
		sortDir = (*common.SortDirection)(params.SortDir)
	}

	paginationDetails := pagefilter.GetPaginatorDetails(params.Limit, params.LastVal, params.LastId, params.SortBy, sortDir)

	// If the limit is not set, remove it from the pagination details.
	if params.Limit == nil {
		paginationDetails.RemoveLimit()
	}

	filts, err := s.getPlayersFilters(params.Name)
	if err != nil {
		l.Error("Failed to parse filters", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to parse filters", err)
		return
	}

	players, err := s.r.GetPlayers(paginationDetails, filts)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrPlayerNotFound):
			players = &pagefilter.PaginatedResponse[models.Player]{
				Items: make([]*models.Player, 0),
				Total: 0,
			}
		default:
			l.Error("Error getting players", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting players", err)
			return
		}
	}

	respArray := make([]api.Player, len(players.Items))
	for i, player := range players.Items {
		respArray[i] = *modelAsApiPlayer(player)
	}

	resp := &api.PlayersResponse{
		Players: respArray,
		Total:   players.Total,
	}

	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}

func modelAsApiPlayer(player *models.Player) *api.Player {
	return &api.Player{
		Id:          utils.Ptr(int64(player.Id)),
		FirstName:   utils.Ptr(player.FirstName),
		LastName:    utils.Ptr(player.LastName),
		Email:       utils.Ptr(openapi_types.Email(player.Email)),
		DateOfBirth: utils.Ptr(openapi_types.Date{Time: player.Dob}),
	}
}

func (s *service) getPlayersFilters(
	name *api.QueryName,
) (*repo.GetPlayersFilters, error) {
	filters := new(repo.GetPlayersFilters)

	if name != nil && *name != "" {
		filters.Name = *name
	}

	return filters, nil
}

func (s *service) CreatePlayer(w http.ResponseWriter, r *http.Request, body0 *api.CreatePlayerJSONBody) {
	l := logging.LoggerFromRequest(r)

	if err := s.validatePlayer(body0); err != nil {
		l.Error("Invalid player", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "invalid player", err)
		return
	}

	p := mapAPIPlayerToModel(body0)

	if err := s.r.CreatePlayer(p); err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicatePlayer):
			l.Error("Player already exists", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusConflict, "player already exists", err)
			return
		default:
			l.Error("Error creating player", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error creating player", err)
			return
		}
	}

	resp := modelAsApiPlayer(p)
	if err := uhttp.Encode(w, http.StatusCreated, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}

func (s *service) validatePlayer(player *api.Player) error {
	if player == nil {
		return errors.New("player is nil")
	}

	if player.FirstName == nil || *player.FirstName == "" {
		return errors.New("first name is required")
	}

	if player.LastName == nil || *player.LastName == "" {
		return errors.New("last name is required")
	}

	if player.Email == nil || *player.Email == "" {
		return errors.New("email is required")
	}

	if player.DateOfBirth == nil {
		return errors.New("date of birth is required")
	}

	return nil
}

func mapAPIPlayerToModel(player *api.Player) *models.Player {
	p := new(models.Player)

	if player.FirstName != nil {
		p.FirstName = *player.FirstName
	}

	if player.LastName != nil {
		p.LastName = *player.LastName
	}

	if player.Email != nil {
		p.Email = string(*player.Email)
	}

	if player.DateOfBirth != nil {
		p.Dob = player.DateOfBirth.Time
	}

	return p
}

func (s *service) GetPlayerByID(w http.ResponseWriter, r *http.Request, id int64) {
	l := logging.LoggerFromRequest(r)

	player, err := s.r.GetPlayer(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrPlayerNotFound):
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "player not found", err)
		default:
			l.Error("Error getting player", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting player", err)
		}
		return
	}

	resp := modelAsApiPlayer(player)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}

func (s *service) UpdatePlayer(w http.ResponseWriter, r *http.Request, id int64, body0 *api.UpdatePlayerJSONBody) {
	l := logging.LoggerFromRequest(r)

	currentPlayer, err := s.r.GetPlayer(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrPlayerNotFound):
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "player not found", err)
		default:
			l.Error("Error getting player", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting player", err)
		}
		return
	}

	changesModel := mapAPIPlayerToModel(body0)
	currentPlayerCopy := *currentPlayer

	// Load the changes into the current player
	if err := patcher.LoadDiff(currentPlayer, changesModel); err != nil {
		l.Error("Failed to load diff", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to load diff", err)
		return
	}

	// See if there are any changes
	if reflect.DeepEqual(currentPlayer, &currentPlayerCopy) {
		l.Info("No changes detected")
		if err := uhttp.Encode(w, http.StatusOK, modelAsApiPlayer(currentPlayer)); err != nil {
			l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
			return
		}
		return
	}

	if err := s.r.UpdatePlayer(id, currentPlayer); err != nil {
		l.Error("Failed to update player", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to update player", err)
		return
	}

	resp := modelAsApiPlayer(currentPlayer)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}
