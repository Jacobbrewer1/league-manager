package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
)

func (r *repository) GetPlayers(params *api.GetPlayersParams) (*api.PlayersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := r.client.GetPlayersWithResponse(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error getting players: %w", err)
	}

	switch {
	case resp.JSON200 != nil:
		return resp.JSON200, nil
	case resp.JSON400 != nil:
		return nil, fmt.Errorf("bad request: %s", *resp.JSON400.Error)
	case resp.JSON500 != nil:
		return nil, fmt.Errorf("internal server error: %s", *resp.JSON500.Message)
	default:
		return nil, fmt.Errorf("unknown response: %d", resp.StatusCode())
	}
}
