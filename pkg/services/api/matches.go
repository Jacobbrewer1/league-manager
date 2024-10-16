package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/Jacobbrewer1/league-manager/pkg/logging"
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	repo "github.com/Jacobbrewer1/league-manager/pkg/repositories/api"
	"github.com/Jacobbrewer1/pagefilter"
	"github.com/Jacobbrewer1/pagefilter/common"
	"github.com/Jacobbrewer1/uhttp"
)

func (s *service) GetMatches(w http.ResponseWriter, r *http.Request, params api.GetMatchesParams) {
	l := logging.LoggerFromRequest(r)

	sortDir := new(common.SortDirection)
	if params.SortDir != nil {
		sortDir = (*common.SortDirection)(params.SortDir)
	}

	paginationDetails := pagefilter.GetPaginatorDetails(params.Limit, params.LastVal, params.LastId, params.SortBy, sortDir)

	filts, err := s.getMatchesFilters(params.Season, params.Team, params.Date, params.DateMin, params.DateMax)
	if err != nil {
		l.Error("Failed to parse filters", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to parse filters", err)
		return
	}

	matches, err := s.r.GetMatches(paginationDetails, filts)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrMatchNotFound):
			matches = &pagefilter.PaginatedResponse[models.Match]{
				Items: make([]*models.Match, 0),
				Total: 0,
			}
		default:
			l.Error("Error getting matches", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting matches", err)
			return
		}
	}

	if err := uhttp.Encode(w, http.StatusOK, matches); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error encoding response", err)
		return
	}
}

func (s *service) getMatchesFilters(
	season *api.QuerySeason,
	team *api.QueryTeam,
	date *api.QueryDate,
	dateFrom *api.QueryDateMin,
	dateTo *api.QueryDateMax,
) (*repo.GetMatchesFilters, error) {
	f := new(repo.GetMatchesFilters)

	if season != nil {
		f.Season = *season
	}

	if team != nil {
		f.Team = *team
	}

	if date != nil {
		f.Date = date
	}

	if dateFrom != nil {
		f.DateFrom = dateFrom
	}

	if dateTo != nil {
		f.DateTo = dateTo
	}

	return f, nil
}

func (s *service) CreateMatch(w http.ResponseWriter, r *http.Request, body0 *api.CreateMatchJSONBody) {
	//TODO implement me
	panic("implement me")
}
