package api

import (
	"fmt"

	"github.com/Jacobbrewer1/league-manager/pkg/models"
)

func (r *repository) GetScoreByMatchAndPartnership(matchID, partnershipID int64) (*models.Score, error) {
	sqlStmt := `
		SELECT id
		FROM score s
		WHERE s.match_id = ?
		AND s.partnership_id = ?
	`

	var scoreID int64
	if err := r.db.Get(&scoreID, sqlStmt, matchID, partnershipID); err != nil {
		return nil, fmt.Errorf("get score by match and partnership: %w", err)
	}

	score, err := models.ScoreById(r.db, int(scoreID))
	if err != nil {
		return nil, fmt.Errorf("get score by id: %w", err)
	}

	return score, nil
}

func (r *repository) SaveScore(score *models.Score) error {
	if err := score.Save(r.db); err != nil {
		return fmt.Errorf("error saving score: %w", err)
	}

	return nil
}
