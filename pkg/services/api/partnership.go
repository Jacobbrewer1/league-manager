package api

import (
	"fmt"

	"github.com/Jacobbrewer1/league-manager/pkg/codegen/apis/api"
	"github.com/Jacobbrewer1/league-manager/pkg/models"
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
