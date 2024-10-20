package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/jacobbrewer1/goschema/pkg/usql"
	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/jacobbrewer1/league-manager/pkg/logging"
	"github.com/jacobbrewer1/league-manager/pkg/models"
	repo "github.com/jacobbrewer1/league-manager/pkg/repositories/api"
	"github.com/jacobbrewer1/league-manager/pkg/utils"
	"github.com/jacobbrewer1/pagefilter"
	"github.com/jacobbrewer1/pagefilter/common"
	"github.com/jacobbrewer1/uhttp"
)

func (s *service) GetGames(w http.ResponseWriter, r *http.Request, params api.GetGamesParams) {
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

	matches, err := s.r.GetGames(paginationDetails, filts)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrMatchNotFound):
			matches = &pagefilter.PaginatedResponse[models.Game]{
				Items: make([]*models.Game, 0),
				Total: 0,
			}
		default:
			l.Error("Error getting matches", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting matches", err)
			return
		}
	}

	respMatches := make([]api.Game, 0, len(matches.Items))
	for _, m := range matches.Items {
		respMatch := new(api.Game)
		respMatch.Id = utils.Ptr(int64(m.Id))
		respMatch.MatchDate = utils.Ptr(m.MatchDate.Format(time.RFC3339))

		winningTeam := api.WinningTeam_home
		if string(m.WinningTeam) == models.GameWinningTeamAway {
			winningTeam = api.WinningTeam_away
		}
		respMatch.WinningTeam = &winningTeam

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

		respMatch.HomeTeam = &api.GamePartnership{
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

		respMatch.AwayTeam = &api.GamePartnership{
			Partnership: awayAPI,
			Scores:      modelAsAPIScore(awayScore),
		}

		respMatches = append(respMatches, *respMatch)
	}

	resp := &api.GamesResponse{
		Games: respMatches,
		Total: matches.Total,
	}

	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error encoding response", err)
		return
	}
}

func modelAsAPIScore(score *models.Score) *api.Scores {
	s := &api.Scores{
		FirstSet:  int64(score.FirstSetScore),
		SecondSet: int64(score.SecondSetScore),
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

func (s *service) CreateGame(w http.ResponseWriter, r *http.Request, body0 *api.CreateGameJSONBody) {
	l := logging.LoggerFromRequest(r)

	if err := s.validateCreateMatch(body0); err != nil {
		l.Error("Failed to validate request", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to validate request", err)
		return
	}

	match, err := matchAsModel(body0)
	if err != nil {
		l.Error("Failed to convert request to model", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to convert request to model", err)
		return
	}

	homePartnership, err := s.partnershipAsModel(body0.HomeTeam.Partnership)
	if err != nil {
		l.Error("Failed to get home partnership", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to get home partnership", err)
		return
	}

	if err := s.r.SavePartnership(homePartnership); err != nil {
		l.Error("Failed to save home partnership", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to save home partnership", err)
		return
	}

	awayPartnership, err := s.partnershipAsModel(body0.AwayTeam.Partnership)
	if err != nil {
		l.Error("Failed to get away partnership", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to get away partnership", err)
		return
	}

	if err := s.r.SavePartnership(awayPartnership); err != nil {
		l.Error("Failed to save away partnership", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to save away partnership", err)
		return
	}

	match.HomePartnersId = homePartnership.Id
	match.AwayPartnersId = awayPartnership.Id
	if err := s.r.CreateMatch(match); err != nil {
		l.Error("Failed to create match", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to create match", err)
		return
	}

	homeScore := scoreAsModel(body0.HomeTeam.Scores)
	if err := s.handleScore(homeScore, homePartnership.Id, match.Id); err != nil {
		l.Error("Failed to handle home score", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to handle home score", err)
		return
	}

	awayScore := scoreAsModel(body0.AwayTeam.Scores)
	if err := s.handleScore(awayScore, awayPartnership.Id, match.Id); err != nil {
		l.Error("Failed to handle away score", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to handle away score", err)
		return
	}

	uhttp.SendMessageWithStatus(w, http.StatusCreated, "match created")
}

func (s *service) handleScore(score *models.Score, partnershipID int, matchID int) error {
	score.PartnershipId = partnershipID
	score.GameId = matchID

	if err := s.r.SaveScore(score); err != nil {
		return fmt.Errorf("failed to save score: %w", err)
	}

	return nil
}

func scoreAsModel(s *api.Scores) *models.Score {
	score := new(models.Score)

	score.FirstSetScore = int(s.FirstSet)
	score.SecondSetScore = int(s.SecondSet)

	if s.ThirdSet != nil {
		score.ThirdSetScore = *usql.NewNullInt64(*s.ThirdSet)
	}

	return score
}

func matchAsModel(m *api.CreateGameJSONBody) (*models.Game, error) {
	match := new(models.Game)

	date, err := time.Parse(time.RFC3339, *m.MatchDate)
	if err != nil {
		return nil, fmt.Errorf("invalid match date: %w", err)
	}
	match.MatchDate = date

	if m.Season == nil || m.Season.Id == nil {
		return nil, errors.New("season is required")
	}
	match.SeasonId = int(*m.Season.Id)

	switch *m.WinningTeam {
	case api.WinningTeam_home:
		match.WinningTeam = usql.NewEnum(models.GameWinningTeamHome)
	case api.WinningTeam_away:
		match.WinningTeam = usql.NewEnum(models.GameWinningTeamAway)
	}

	return match, nil
}

func (s *service) validateCreateMatch(body *api.CreateGameJSONBody) error {
	if body == nil {
		return errors.New("request body is nil")
	}

	if body.Season == nil {
		return errors.New("season is required")
	} else if body.Season.Id == nil {
		return errors.New("season id is required")
	}

	if body.MatchDate == nil {
		return errors.New("match date is required")
	} else if _, err := time.Parse(time.RFC3339, *body.MatchDate); err != nil {
		return errors.New("invalid match date")
	}

	if body.HomeTeam == nil {
		return errors.New("home team is required")
	} else if body.HomeTeam.Partnership == nil {
		return errors.New("home team partnership is required")
	} else if body.AwayTeam.Partnership.Team == nil {
		return errors.New("home team team is required")
	} else if body.AwayTeam.Partnership.Team.Id == nil {
		return errors.New("home team team id is required")
	}

	if body.HomeTeam.Partnership.PlayerA == nil {
		return errors.New("home team player A is required")
	} else if body.HomeTeam.Partnership.PlayerA.Id == nil {
		return errors.New("home team player A id is required")
	}

	if body.HomeTeam.Partnership.PlayerB == nil {
		return errors.New("home team player B is required")
	} else if body.HomeTeam.Partnership.PlayerB.Id == nil {
		return errors.New("home team player B id is required")
	}

	if body.HomeTeam.Scores == nil {
		return errors.New("home team scores is required")
	}

	if body.AwayTeam == nil {
		return errors.New("home team is required")
	} else if body.AwayTeam.Partnership == nil {
		return errors.New("home team partnership is required")
	} else if body.AwayTeam.Partnership.Team == nil {
		return errors.New("home team team is required")
	} else if body.AwayTeam.Partnership.Team.Id == nil {
		return errors.New("home team team id is required")
	}

	if body.AwayTeam.Partnership.PlayerA == nil {
		return errors.New("home team player A is required")
	} else if body.AwayTeam.Partnership.PlayerA.Id == nil {
		return errors.New("home team player A id is required")
	}

	if body.AwayTeam.Partnership.PlayerB == nil {
		return errors.New("home team player B is required")
	} else if body.AwayTeam.Partnership.PlayerB.Id == nil {
		return errors.New("home team player B id is required")
	}

	if body.AwayTeam.Scores == nil {
		return errors.New("home team scores is required")
	}

	return nil
}
