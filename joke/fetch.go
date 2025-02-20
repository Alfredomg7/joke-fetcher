package joke

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JokeResponse represents the joke API response structure
type JokeResponse struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// JokeFetcher defines the interface for fetching jokes
type JokeFetcher interface {
	FetchJoke() (*JokeResponse, error)
}

// APIJokeFetcher is the real implementation that fetches jokes from an API
type APIJokeFetcher struct{}

// FetchJoke fetches a joke from the API and returns a pointer to the struct
func (a APIJokeFetcher) FetchJoke() (*JokeResponse, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/jokes/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch joke, received non-200 response")
	}

	var joke JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return nil, err
	}

	return &joke, nil // Returning pointer to avoid struct copying
}

// Mock implementation of JokeFetcher for testing
type MockJokeFetcher struct {
	ShouldFail bool
}

func (m MockJokeFetcher) FetchJoke() (*JokeResponse, error) {
	if m.ShouldFail {
		return nil, errors.New("mock fetcher failed")
	}
	return &JokeResponse{
		Setup:     "Why did the chicken cross the road?",
		Punchline: "To get to the other side!",
	}, nil
}
