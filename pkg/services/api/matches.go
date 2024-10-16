package api

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/Jacobbrewer1/league-manager/pkg/logging"
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	repo "github.com/Jacobbrewer1/league-manager/pkg/repositories/api"
	"github.com/Jacobbrewer1/league-manager/pkg/utils"
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

	respMatches := make([]api.Match, 0, len(matches.Items))
	for _, m := range matches.Items {
		respMatch := new(api.Match)
		respMatch.Id = utils.Ptr(int64(m.Id))
		respMatch.MatchDate = utils.Ptr(m.MatchDate.Format(time.RFC3339))

		season, err := s.r.GetSeason(int64(m.SeasonId))
		if err != nil {
			l.Error("Error getting season", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting season", err)
			return
		}
		respMatch.Season = modelAsAPISeason(season)

		homePartnership, err := s.r.GetPartnership(int64(m.HomePartnersId))
		if err != nil {
			l.Error("Error getting home partnership", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting home partnership", err)
			return
		}

		homeAPI, err := s.modelAsAPIPartnership(homePartnership)
		if err != nil {
			l.Error("Error getting home partnership", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting home partnership", err)
			return
		}

		homeScore, err := s.r.GetScoreByMatchAndPartnership(int64(m.Id), int64(m.HomePartnersId))
		if err != nil {
			l.Error("Error getting home score", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting home score", err)
			return
		}

		respMatch.HomeTeam = &api.MatchPartnership{
			Partnership: homeAPI,
			Scores:      modelAsAPIScore(homeScore),
		}

		awayPartnership, err := s.r.GetPartnership(int64(m.AwayPartnersId))
		if err != nil {
			l.Error("Error getting away partnership", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting away partnership", err)
			return
		}

		awayAPI, err := s.modelAsAPIPartnership(awayPartnership)
		if err != nil {
			l.Error("Error getting away partnership", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting away partnership", err)
			return
		}

		awayScore, err := s.r.GetScoreByMatchAndPartnership(int64(m.Id), int64(m.AwayPartnersId))
		if err != nil {
			l.Error("Error getting away score", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting away score", err)
			return
		}

		respMatch.AwayTeam = &api.MatchPartnership{
			Partnership: awayAPI,
			Scores:      modelAsAPIScore(awayScore),
		}

		respMatches = append(respMatches, *respMatch)
	}

	resp := &api.MatchesResponse{
		Matches: respMatches,
		Total:   matches.Total,
	}

	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error encoding response", err)
		return
	}
}

func modelAsAPIScore(score *models.Score) *api.Scores {
	s := &api.Scores{
		FirstSet:  utils.Ptr(int64(score.FirstSetScore)),
		SecondSet: utils.Ptr(int64(score.SecondSetScore)),
	}

	if score.ThirdSetScore.Valid {
		s.ThirdSet = utils.Ptr(score.ThirdSetScore.Int64)
	}

	return s
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
