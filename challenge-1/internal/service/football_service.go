package service

import (
	"fmt"
	"strconv"

	client "github.com/wahyudin/football-goals-calculator/internal/clients"
)

type FootballService struct {
	apiClient *client.APIClient
}

func NewFootballService() *FootballService {
	return &FootballService{
		apiClient: client.NewAPIClient(),
	}
}

func (s *FootballService) CalculateTotalGoals(teamName string, year int) (int, error) {
	totalGoals := 0

	// Hitung gol ketika tim bermain sebagai team1
	team1Goals, err := s.getGoalsAsTeam1(teamName, year)
	if err != nil {
		return 0, fmt.Errorf("failed to get goals as team1: %w", err)
	}

	// Hitung gol ketika tim bermain sebagai team2
	team2Goals, err := s.getGoalsAsTeam2(teamName, year)
	if err != nil {
		return 0, fmt.Errorf("failed to get goals as team2: %w", err)
	}

	totalGoals = team1Goals + team2Goals
	return totalGoals, nil
}

func (s *FootballService) getGoalsAsTeam1(teamName string, year int) (int, error) {
	totalGoals := 0
	page := 1

	for {
		response, err := s.apiClient.FetchMatches(teamName, year, "team1", page)
		if err != nil {
			return 0, err
		}

		// Kalkulasi gol dari halaman ini
		for _, match := range response.Data {
			goals, err := strconv.Atoi(match.Team1Goals)
			if err != nil {
				fmt.Printf("Warning: Invalid team1goals value '%s' for match %s vs %s\n",
					match.Team1Goals, match.Team1, match.Team2)
				continue
			}
			totalGoals += goals
		}

		// Cek apakah sudah mencapai halaman terakhir
		if page >= response.TotalPages {
			break
		}
		page++
	}

	return totalGoals, nil
}

func (s *FootballService) getGoalsAsTeam2(teamName string, year int) (int, error) {
	totalGoals := 0
	page := 1

	for {
		response, err := s.apiClient.FetchMatches(teamName, year, "team2", page)
		if err != nil {
			return 0, err
		}

		// Kalkulasi gol dari halaman ini
		for _, match := range response.Data {
			goals, err := strconv.Atoi(match.Team2Goals)
			if err != nil {
				fmt.Printf("Warning: Invalid team2goals value '%s' for match %s vs %s\n",
					match.Team2Goals, match.Team1, match.Team2)
				continue
			}
			totalGoals += goals
		}

		// Cek apakah sudah mencapai halaman terakhir
		if page >= response.TotalPages {
			break
		}
		page++
	}

	return totalGoals, nil
}
