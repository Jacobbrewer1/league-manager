// Code generated by mockery. DO NOT EDIT.

package api

import (
	models "github.com/Jacobbrewer1/league-manager/pkg/models"
	pagefilter "github.com/Jacobbrewer1/pagefilter"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CreateMatch provides a mock function with given fields: match
func (_m *MockRepository) CreateMatch(match *models.Game) error {
	ret := _m.Called(match)

	if len(ret) == 0 {
		panic("no return value specified for CreateMatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Game) error); ok {
		r0 = rf(match)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePlayer provides a mock function with given fields: player
func (_m *MockRepository) CreatePlayer(player *models.Player) error {
	ret := _m.Called(player)

	if len(ret) == 0 {
		panic("no return value specified for CreatePlayer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Player) error); ok {
		r0 = rf(player)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSeason provides a mock function with given fields: season
func (_m *MockRepository) CreateSeason(season *models.Season) error {
	ret := _m.Called(season)

	if len(ret) == 0 {
		panic("no return value specified for CreateSeason")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Season) error); ok {
		r0 = rf(season)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTeam provides a mock function with given fields: team
func (_m *MockRepository) CreateTeam(team *models.Team) error {
	ret := _m.Called(team)

	if len(ret) == 0 {
		panic("no return value specified for CreateTeam")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Team) error); ok {
		r0 = rf(team)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGames provides a mock function with given fields: details, filters
func (_m *MockRepository) GetGames(details *pagefilter.PaginatorDetails, filters *GetMatchesFilters) (*pagefilter.PaginatedResponse[models.Game], error) {
	ret := _m.Called(details, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetGames")
	}

	var r0 *pagefilter.PaginatedResponse[models.Game]
	var r1 error
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetMatchesFilters) (*pagefilter.PaginatedResponse[models.Game], error)); ok {
		return rf(details, filters)
	}
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetMatchesFilters) *pagefilter.PaginatedResponse[models.Game]); ok {
		r0 = rf(details, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagefilter.PaginatedResponse[models.Game])
		}
	}

	if rf, ok := ret.Get(1).(func(*pagefilter.PaginatorDetails, *GetMatchesFilters) error); ok {
		r1 = rf(details, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPartnership provides a mock function with given fields: id
func (_m *MockRepository) GetPartnership(id int64) (*models.Partnership, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPartnership")
	}

	var r0 *models.Partnership
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Partnership, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Partnership); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Partnership)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayer provides a mock function with given fields: id
func (_m *MockRepository) GetPlayer(id int64) (*models.Player, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPlayer")
	}

	var r0 *models.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Player, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Player); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayers provides a mock function with given fields: details, filters
func (_m *MockRepository) GetPlayers(details *pagefilter.PaginatorDetails, filters *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error) {
	ret := _m.Called(details, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetPlayers")
	}

	var r0 *pagefilter.PaginatedResponse[models.Player]
	var r1 error
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetPlayersFilters) (*pagefilter.PaginatedResponse[models.Player], error)); ok {
		return rf(details, filters)
	}
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetPlayersFilters) *pagefilter.PaginatedResponse[models.Player]); ok {
		r0 = rf(details, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagefilter.PaginatedResponse[models.Player])
		}
	}

	if rf, ok := ret.Get(1).(func(*pagefilter.PaginatorDetails, *GetPlayersFilters) error); ok {
		r1 = rf(details, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScoreByMatchAndPartnership provides a mock function with given fields: matchID, partnershipID
func (_m *MockRepository) GetScoreByMatchAndPartnership(matchID int64, partnershipID int64) (*models.Score, error) {
	ret := _m.Called(matchID, partnershipID)

	if len(ret) == 0 {
		panic("no return value specified for GetScoreByMatchAndPartnership")
	}

	var r0 *models.Score
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (*models.Score, error)); ok {
		return rf(matchID, partnershipID)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) *models.Score); ok {
		r0 = rf(matchID, partnershipID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Score)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(matchID, partnershipID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSeason provides a mock function with given fields: id
func (_m *MockRepository) GetSeason(id int64) (*models.Season, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetSeason")
	}

	var r0 *models.Season
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Season, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Season); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Season)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSeasons provides a mock function with given fields: details, filters
func (_m *MockRepository) GetSeasons(details *pagefilter.PaginatorDetails, filters *GetSeasonsFilters) (*pagefilter.PaginatedResponse[models.Season], error) {
	ret := _m.Called(details, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetSeasons")
	}

	var r0 *pagefilter.PaginatedResponse[models.Season]
	var r1 error
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetSeasonsFilters) (*pagefilter.PaginatedResponse[models.Season], error)); ok {
		return rf(details, filters)
	}
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetSeasonsFilters) *pagefilter.PaginatedResponse[models.Season]); ok {
		r0 = rf(details, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagefilter.PaginatedResponse[models.Season])
		}
	}

	if rf, ok := ret.Get(1).(func(*pagefilter.PaginatorDetails, *GetSeasonsFilters) error); ok {
		r1 = rf(details, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeam provides a mock function with given fields: id
func (_m *MockRepository) GetTeam(id int64) (*models.Team, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTeam")
	}

	var r0 *models.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Team, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Team); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeams provides a mock function with given fields: details, filters
func (_m *MockRepository) GetTeams(details *pagefilter.PaginatorDetails, filters *GetTeamsFilters) (*pagefilter.PaginatedResponse[models.Team], error) {
	ret := _m.Called(details, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetTeams")
	}

	var r0 *pagefilter.PaginatedResponse[models.Team]
	var r1 error
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetTeamsFilters) (*pagefilter.PaginatedResponse[models.Team], error)); ok {
		return rf(details, filters)
	}
	if rf, ok := ret.Get(0).(func(*pagefilter.PaginatorDetails, *GetTeamsFilters) *pagefilter.PaginatedResponse[models.Team]); ok {
		r0 = rf(details, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pagefilter.PaginatedResponse[models.Team])
		}
	}

	if rf, ok := ret.Get(1).(func(*pagefilter.PaginatorDetails, *GetTeamsFilters) error); ok {
		r1 = rf(details, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePartnership provides a mock function with given fields: partnership
func (_m *MockRepository) SavePartnership(partnership *models.Partnership) error {
	ret := _m.Called(partnership)

	if len(ret) == 0 {
		panic("no return value specified for SavePartnership")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Partnership) error); ok {
		r0 = rf(partnership)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveScore provides a mock function with given fields: score
func (_m *MockRepository) SaveScore(score *models.Score) error {
	ret := _m.Called(score)

	if len(ret) == 0 {
		panic("no return value specified for SaveScore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Score) error); ok {
		r0 = rf(score)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePlayer provides a mock function with given fields: id, player
func (_m *MockRepository) UpdatePlayer(id int64, player *models.Player) error {
	ret := _m.Called(id, player)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePlayer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *models.Player) error); ok {
		r0 = rf(id, player)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSeason provides a mock function with given fields: id, season
func (_m *MockRepository) UpdateSeason(id int64, season *models.Season) error {
	ret := _m.Called(id, season)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSeason")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *models.Season) error); ok {
		r0 = rf(id, season)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTeam provides a mock function with given fields: id, team
func (_m *MockRepository) UpdateTeam(id int64, team *models.Team) error {
	ret := _m.Called(id, team)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTeam")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *models.Team) error); ok {
		r0 = rf(id, team)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
