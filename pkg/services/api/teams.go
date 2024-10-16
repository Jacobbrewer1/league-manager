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

func (s *service) GetTeams(w http.ResponseWriter, r *http.Request, params api.GetTeamsParams) {
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

	filts, err := s.getTeamsFilters(params.Name)
	if err != nil {
		l.Error("Failed to parse filters", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to parse filters", err)
		return
	}

	teams, err := s.r.GetTeams(paginationDetails, filts)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrTeamNotFound):
			teams = &pagefilter.PaginatedResponse[models.Team]{
				Items: make([]*models.Team, 0),
				Total: 0,
			}
		default:
			l.Error("Error getting teams", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting teams", err)
			return
		}
	}

	respArray := make([]api.Team, len(teams.Items))
	for i, team := range teams.Items {
		respArray[i] = *modelAsApiTeam(team)
	}

	resp := &api.TeamsResponse{
		Teams: respArray,
		Total: teams.Total,
	}

	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error encoding response", err)
		return
	}
}

func modelAsApiTeam(team *models.Team) *api.Team {
	return &api.Team{
		ContactEmail: utils.Ptr(openapi_types.Email(team.ContactEmail)),
		ContactPhone: utils.Ptr(team.ContactMobile),
		Id:           utils.Ptr(int64(team.Id)),
		Name:         utils.Ptr(team.Name),
	}
}

func (s *service) getTeamsFilters(
	name *api.QueryName,
) (*repo.GetTeamsFilters, error) {
	filters := new(repo.GetTeamsFilters)

	if name != nil {
		filters.Name = *name
	}

	return filters, nil
}

func (s *service) CreateTeam(w http.ResponseWriter, r *http.Request, body0 *api.CreateTeamJSONBody) {
	//TODO implement me
	panic("implement me")
}
