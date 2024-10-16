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
)

func (s *service) GetSeasons(w http.ResponseWriter, r *http.Request, params api.GetSeasonsParams) {
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

	filts, err := s.getSeasonsFilters(params.Name)
	if err != nil {
		l.Error("Failed to parse filters", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to parse filters", err)
		return
	}

	seasons, err := s.r.GetSeasons(paginationDetails, filts)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrSeasonNotFound):
			seasons = &pagefilter.PaginatedResponse[models.Season]{
				Items: make([]*models.Season, 0),
				Total: 0,
			}
		default:
			l.Error("Error getting seasons", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting seasons", err)
			return
		}
	}

	respArray := make([]api.Season, len(seasons.Items))
	for i, season := range seasons.Items {
		respArray[i] = *modelAsAPISeason(season)
	}

	resp := &api.SeasonsResponse{
		Seasons: respArray,
		Total:   seasons.Total,
	}

	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Error encoding response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error encoding response", err)
	}
}

func (s *service) getSeasonsFilters(
	name *string,
) (*repo.GetSeasonsFilters, error) {
	filters := new(repo.GetSeasonsFilters)

	if name != nil {
		filters.Name = *name
	}

	return filters, nil
}

func (s *service) CreateSeason(w http.ResponseWriter, r *http.Request, body0 *api.CreateSeasonJSONBody) {
	l := logging.LoggerFromRequest(r)

	if err := s.validateSeason(body0); err != nil {
		l.Error("Failed to validate season", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to validate season", err)
		return
	}

	season := mapAPISeasonToModel(body0)

	if err := s.r.CreateSeason(season); err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicateSeason):
			l.Error("Season already exists", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusConflict, "season already exists", err)
			return
		default:
			l.Error("Error creating season", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error creating season", err)
			return
		}
	}

	resp := modelAsAPISeason(season)
	if err := uhttp.Encode(w, http.StatusCreated, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to encode response", err)
	}
}

func modelAsAPISeason(season *models.Season) *api.Season {
	return &api.Season{
		Id:   utils.Ptr(int64(season.Id)),
		Name: utils.Ptr(season.Name),
	}
}

func mapAPISeasonToModel(season *api.CreateSeasonJSONBody) *models.Season {
	return &models.Season{
		Name: *season.Name,
	}
}

func (s *service) validateSeason(
	season *api.CreateSeasonJSONBody,
) error {
	if season == nil {
		return errors.New("body is required")
	}

	if season.Name == nil {
		return errors.New("name is required")
	}

	return nil
}

func (s *service) GetSeasonByID(w http.ResponseWriter, r *http.Request, id int64) {
	l := logging.LoggerFromRequest(r)

	season, err := s.r.GetSeason(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrSeasonNotFound):
			l.Error("Season not found", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "season not found", err)
			return
		default:
			l.Error("Error getting season", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting season", err)
			return
		}
	}

	resp := modelAsAPISeason(season)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to encode response", err)
	}
}

func (s *service) UpdateSeason(w http.ResponseWriter, r *http.Request, id int64, body0 *api.UpdateSeasonJSONBody) {
	l := logging.LoggerFromRequest(r)

	currentSeason, err := s.r.GetSeason(id)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrSeasonNotFound):
			l.Error("Season not found", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusNotFound, "season not found", err)
			return
		default:
			l.Error("Error getting season", slog.String(logging.KeyError, err.Error()))
			uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error getting season", err)
			return
		}
	}

	changesModel := mapAPISeasonToModel(body0)
	currentSeasonCopy := *currentSeason

	if err := patcher.LoadDiff(currentSeason, changesModel); err != nil {
		l.Error("Failed to load diff", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusBadRequest, "failed to load diff", err)
		return
	}

	if reflect.DeepEqual(currentSeason, &currentSeasonCopy) {
		l.Info("No changes detected")
		if err := uhttp.Encode(w, http.StatusOK, modelAsAPISeason(currentSeason)); err != nil {
			l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
			return
		}
		return
	}

	if err := s.r.UpdateSeason(id, currentSeason); err != nil {
		l.Error("Error updating season", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "error updating season", err)
		return
	}

	resp := modelAsAPISeason(currentSeason)
	if err := uhttp.Encode(w, http.StatusOK, resp); err != nil {
		l.Error("Failed to encode response", slog.String(logging.KeyError, err.Error()))
		uhttp.SendErrorMessageWithStatus(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}
