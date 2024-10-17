package api

import (
	"errors"
	"fmt"

	"github.com/Jacobbrewer1/league-manager/pkg/models"
)

func (r *repository) GetPartnership(id int64) (*models.Partnership, error) {
	p, err := models.PartnershipById(r.db, int(id))
	if err != nil {
		return nil, fmt.Errorf("error getting partnership: %w", err)
	}

	return p, nil
}

func (r *repository) SavePartnership(partnership *models.Partnership) error {
	// Ensure that player A has the lower ID and player B has the higher ID
	if partnership.PlayerAId > partnership.PlayerBId {
		partnership.PlayerAId, partnership.PlayerBId = partnership.PlayerBId, partnership.PlayerAId
	}

	// See if the partnership already exists
	existingPartnership, err := models.PartnershipByPlayerAIdPlayerBIdTeamId(r.db, partnership.PlayerAId, partnership.PlayerBId, partnership.TeamId)
	if err != nil {
		return fmt.Errorf("error getting partnership: %w", err)
	} else if existingPartnership != nil {
		partnership.Id = existingPartnership.Id
	}

	if err := partnership.Save(r.db); err != nil && !errors.Is(err, models.ErrNoAffectedRows) {
		return fmt.Errorf("error saving partnership: %w", err)
	}

	return nil
}
