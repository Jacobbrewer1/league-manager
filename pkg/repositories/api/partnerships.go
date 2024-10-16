package api

import (
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
