package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/Jacobbrewer1/league-manager/pkg/logging"
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	repo "github.com/Jacobbrewer1/league-manager/pkg/repositories/api"
	"github.com/Jacobbrewer1/league-manager/pkg/utils"
	"github.com/Jacobbrewer1/pagefilter"
	"github.com/Jacobbrewer1/pagefilter/common"
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
		respArray[i] = *s.modelAsApiPlayer(player)
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

func (s *service) modelAsApiPlayer(player *models.Player) *api.Player {
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
	//TODO implement me
	panic("implement me")
}
