package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wahyudin/football-goals-calculator/internal/models"
)

type APIClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewAPIClient() *APIClient {
	return &APIClient{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    "https://jsonmock.hackerrank.com/api",
	}
}

func (c *APIClient) FetchMatches(teamName string, year int, teamPosition string, page int) (*models.APIResponse, error) {
	url := fmt.Sprintf("%s/football_matches?year=%d&%s=%s&page=%d",
		c.baseURL, year, teamPosition, teamName, page)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var apiResponse models.APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	return &apiResponse, nil
}
