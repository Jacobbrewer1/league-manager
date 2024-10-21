package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
	cliRepo "github.com/jacobbrewer1/league-manager/pkg/repositories/cli"
	"github.com/jacobbrewer1/league-manager/pkg/services/cli"
)

const (
	// EnvApiHost is the environment variable for the API host.
	EnvApiHost = "LEAGUE_MANAGER_API_HOST"
)

func main() {
	host := os.Getenv(EnvApiHost)
	if host == "" {
		if err := huh.NewInput().Title("Please enter the API Host").Value(&host).Validate(func(s string) error {
			_, err := parseURL(s)
			return err
		}).Run(); err != nil {
			panic(fmt.Errorf("error getting API host: %w", err))
		}
	}

	u, err := parseURL(host)
	if err != nil {
		panic(fmt.Errorf("error parsing URL: %w", err))
	}
	host = u

	client, err := api.NewClientWithResponses(host)
	if err != nil {
		panic(fmt.Errorf("error creating client: %w", err))
	}

	cliRepository := cliRepo.NewRepository(client)
	if err := cli.NewService(cliRepository).Run(); err != nil {
		panic(fmt.Errorf("error running service: %w", err))
	}
}

func parseURL(s string) (string, error) {
	u, err := url.Parse(strings.TrimSuffix(strings.TrimPrefix(s, "\""), "\""))
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	return u.String(), nil
}
