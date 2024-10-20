package api

import (
	"errors"
	"fmt"

	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/jacobbrewer1/league-manager/pkg/models"
)

func (s *service) modelAsAPIPartnership(m *models.Partnership) (*api.Partnership, error) {
	playerA, err := s.r.GetPlayer(int64(m.PlayerAId))
	if err != nil {
		return nil, fmt.Errorf("error getting player A: %w", err)
	}

	playerB, err := s.r.GetPlayer(int64(m.PlayerBId))
	if err != nil {
		return nil, fmt.Errorf("error getting player B: %w", err)
	}

	team, err := s.r.GetTeam(int64(m.TeamId))
	if err != nil {
		return nil, fmt.Errorf("error getting team: %w", err)
	}

	return &api.Partnership{
		PlayerA: modelAsApiPlayer(playerA),
		PlayerB: modelAsApiPlayer(playerB),
		Team:    modelAsApiTeam(team),
	}, nil
}

func (s *service) partnershipAsModel(p *api.Partnership) (*models.Partnership, error) {
	if p == nil {
		return nil, errors.New("partnership is nil")
	}

	if p.PlayerA == nil {
		return nil, errors.New("player A is required")
	} else if p.PlayerA.Id == nil {
		return nil, errors.New("player A ID is required")
	}

	playerA, err := s.r.GetPlayer(*p.PlayerA.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting player A: %w", err)
	}

	if p.PlayerB == nil {
		return nil, errors.New("player B is required")
	} else if p.PlayerB.Id == nil {
		return nil, errors.New("player B ID is required")
	}

	playerB, err := s.r.GetPlayer(*p.PlayerB.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting player B: %w", err)
	}

	if p.Team == nil {
		return nil, fmt.Errorf("team is required")
	} else if p.Team.Id == nil {
		return nil, fmt.Errorf("team ID is required")
	}

	team, err := s.r.GetTeam(*p.Team.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting team: %w", err)
	}

	return &models.Partnership{
		PlayerAId: playerA.Id,
		PlayerBId: playerB.Id,
		TeamId:    team.Id,
	}, nil
}
