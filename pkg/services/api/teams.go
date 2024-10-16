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
	l := logging.LoggerFromRequest(r)

	if err := s.validateTeam(body0); err != nil {
		l.Error("Failed to validate team", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to validate team", err)
		return
	}

	team := mapAPITeamToModel(body0)
	err := s.r.CreateTeam(team)
	if err != nil {
		l.Error("Failed to create team", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to create team", err)
		return
	}

	resp := modelAsApiTeam(team)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}

func mapAPITeamToModel(team *api.CreateTeamJSONBody) *models.Team {
	t := new(models.Team)

	if team.ContactEmail != nil {
		t.ContactEmail = string(*team.ContactEmail)
	}

	if team.ContactPhone != nil {
		t.ContactMobile = *team.ContactPhone
	}

	if team.Name != nil {
		t.Name = *team.Name
	}

	return t
}

func (s *service) validateTeam(body0 *api.CreateTeamJSONBody) error {
	if body0 == nil {
		return errors.New("body is nil")
	}

	if body0.Name == nil {
		return errors.New("name is nil")
	}

	if *body0.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

func (s *service) GetTeamByID(w http.ResponseWriter, r *http.Request, id int64) {
	l := logging.LoggerFromRequest(r)

	team, err := s.r.GetTeam(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrTeamNotFound):
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "team not found", err)
			return
		default:
			l.Error("Error getting team", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting team", err)
			return
		}
	}

	resp := modelAsApiTeam(team)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		return
	}
}

func (s *service) UpdateTeam(w http.ResponseWriter, r *http.Request, id int64, body0 *api.UpdateTeamJSONBody) {
	l := logging.LoggerFromRequest(r)

	currentTeam, err := s.r.GetTeam(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrTeamNotFound):
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "team not found", err)
			return
		default:
			l.Error("Error getting team", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting team", err)
			return
		}
	}

	changesModel := mapAPITeamToModel(body0)
	currentTeamCopy := *currentTeam

	if err := patcher.LoadDiff(currentTeam, changesModel); err != nil {
		l.Error("Failed to load diff", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to load diff", err)
		return
	}

	if reflect.DeepEqual(currentTeam, &currentTeamCopy) {
		l.Info("No changes detected")
		if err := uhttp.Encode(w, http.StatusOK, modelAsApiTeam(currentTeam)); err != nil {
			l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
			return
		}
		return
	}

	if err := s.r.UpdateTeam(id, currentTeam); err != nil {
		l.Error("Failed to update team", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to update team", err)
		return
	}

	resp := modelAsApiTeam(currentTeam)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		return
	}
}
