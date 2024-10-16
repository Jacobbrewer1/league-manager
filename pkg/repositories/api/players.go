package api

import (
	"github.com/Jacobbrewer1/league-manager/pkg/models"
	"github.com/Jacobbrewer1/pagefilter"
)

func (r repository) GetPlayers(details pagefilter.PaginatorDetails, filters *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error) {
	//TODO implement me
	panic("implement me")
}
